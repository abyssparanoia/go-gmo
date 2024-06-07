package remittance

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/imdario/mergo"
)

// Client ... gmo pg remittance API client
type Client struct {
	HTTPClient *http.Client
	ShopID     string
	ShopPass   string
	APIHost    string
}

// NewClient ... new client
func NewClient(
	shopID,
	shopPass string,
	sandBox bool) (*Client, error) {

	var apiHost string
	if sandBox {
		apiHost = apiHostSandbox
	} else {
		apiHost = apiHostProduction
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second, // limitation of the request time
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: time.Second, // limitation of the dial time
				}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
					// #nosec G402
					CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_128_GCM_SHA256},
				},
				TLSHandshakeTimeout:   10 * time.Second, // limitation of the TLS handshake time
				ResponseHeaderTimeout: 10 * time.Second, // limitation of the response header time
				IdleConnTimeout:       10 * time.Second, // limitation of the idle connection time
				MaxIdleConns:          100,              // limitation of the max idle connections
				MaxConnsPerHost:       100,              // limitation of the max connections per host
				MaxIdleConnsPerHost:   100,              // limitation of the max idle connections per host
			},
		},
		ShopID:   shopID,
		ShopPass: shopPass,
		APIHost:  apiHost,
	}, nil
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.HTTPClient = httpClient
}

type baseRequestBody struct {
	ShopID   string `json:"Shop_ID"`
	ShopPass string `json:"Shop_Pass"`
}

func (c *Client) do(
	path string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {

	requestBodyMap := map[string]interface{}{}
	requestBodyMap["Shop_ID"] = c.ShopID
	requestBodyMap["Shop_Pass"] = c.ShopPass

	if err := mergo.Map(&requestBodyMap, &body); err != nil {
		return nil, err
	}

	requestBodyBytes, err := json.Marshal(requestBodyMap)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.APIHost, path),
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if contains := bytes.Contains(bodyBytes, []byte("ErrCode")); contains {
		errResp := &ErrorResponse{}
		if err := json.Unmarshal(bodyBytes, errResp); err != nil {
			return nil, err
		}
		return nil, errResp
	}

	if err := json.Unmarshal(bodyBytes, respBody); err != nil {
		return nil, err
	}

	return resp, nil
}
