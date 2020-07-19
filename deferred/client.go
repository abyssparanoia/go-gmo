package deferred

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Client ... gmo deferred payment client
type Client struct {
	HTTPClient       *http.Client
	AuthenticationID string
	ShopCode         string
	ConnectPassword  string
	APIHost          string
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 30,
}

// NewClient ... new client
func NewClient(
	authenticationID,
	shopCode,
	connectPassword string,
	sandBox bool,
) (*Client, error) {
	var apiHost string
	if sandBox {
		apiHost = apiHostSandbox
	} else {
		apiHost = apiHostProduction
	}

	return &Client{
		HTTPClient:       defaultHTTPClient,
		AuthenticationID: authenticationID,
		ShopCode:         shopCode,
		ConnectPassword:  connectPassword,
		APIHost:          apiHost,
	}, nil
}

type baseRequestBody struct {
	AuthenticationID string `xml:"authenticationId"`
	ShopCode         string `xml:"shopCode"`
	ConnectPassword  string `xml:"connectPassword"`
}

type httpMethod string

func (m httpMethod) string() string {
	return string(m)
}

const (
	GET    httpMethod = "GET"
	POST   httpMethod = "POST"
	PUT    httpMethod = "PUT"
	DELETE httpMethod = "DELETE"
	PATCH  httpMethod = "PATCH"
)

// reqPtr can be nil if body is not needed
// resPtr can be nil if client want only status code
// errPtr can be nil if client want only status code when status => 400
func (c *Client) doAndUnmarshalXML(ctx context.Context, method httpMethod, rawURL string, paths []string,
	queries map[string]string, reqPtr, respPtr interface{}) (int, error) {
	select {
	default:
	case <-ctx.Done():
		return 0, ctx.Err()
	}
	base, err := url.Parse(rawURL)
	if err != nil {
		return 0, err
	}
	copied := *base
	copied.Path = path.Join(copied.Path, path.Join(paths...))
	if len(queries) > 0 {
		q := copied.Query()
		for k, v := range queries {
			q.Set(k, v)
		}
		copied.RawQuery = q.Encode()
	}

	reqData, err := xml.Marshal(reqPtr)
	if err != nil {
		return 0, fmt.Errorf("httpclient: xml.Marshal(reqPtr) reqPtr=%v : %w", reqPtr, err)
	}

	req, err := http.NewRequest(method.string(), copied.String(), bytes.NewBuffer(reqData))
	if err != nil {
		return 0, fmt.Errorf("httpclient: http.NewRequest reqPtr=%v : %w", reqPtr, err)
	}
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("httpclient: c.HTTPClient.Do reqPtr=%v : %w", reqPtr, err)
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	if respPtr == nil {
		// dont need to unmarshal just return status code
		return status, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("httpclient: ioutil.ReadAll response=%v: %w", resp, err)
	}
	if err := xml.Unmarshal(body, respPtr); err != nil {
		return 0, fmt.Errorf("httpclient: xml.Unmarshal(data, target) data=%v: %w", string(body), err)
	}
	return status, nil
}
