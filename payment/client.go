package payment

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/abyssparanoia/go-gmo/internal/pkg/shiftjis_transformer"
	"moul.io/http2curl"
)

// Client ... gmo pg payment API client
type Client struct {
	HTTPClient  *http.Client
	SiteID      string
	SitePass    string
	ShopID      string
	ShopPass    string
	APIHost     string
	OpenAPIHost string
	SandBox     bool
	Debug       bool
}

// NewClient ... new client
func NewClient(
	siteID,
	sitePass,
	shopID,
	shopPass string,
	sandBox bool) (*Client, error) {

	var apiHost string
	if sandBox {
		apiHost = apiHostSandbox
	} else {
		apiHost = apiHostProduction
	}

	var openAPIHost string
	if sandBox {
		openAPIHost = apiHostOpenAPISandbox
	} else {
		openAPIHost = apiHostOpenAPIProduction
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second, // limitation of the request time
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second, // limitation of the dial time
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
		SiteID:      siteID,
		SitePass:    sitePass,
		ShopID:      shopID,
		ShopPass:    shopPass,
		APIHost:     apiHost,
		OpenAPIHost: openAPIHost,
		SandBox:     sandBox,
	}, nil
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.HTTPClient = httpClient
}

func (c *Client) SetDebug(debug bool) {
	c.Debug = debug
}

type baseRequestBody struct {
	SiteID   string `schema:"SiteID,omitempty"`
	SitePass string `schema:"SitePass,omitempty"`
	ShopID   string `schema:"ShopID,omitempty"`
	ShopPass string `schema:"ShopPass,omitempty"`
}

func (c *Client) do(
	path string,
	body interface{},
	respBody interface{},
) (*http.Response, error) {

	baseForms := url.Values{}
	err := parser.Encoder().Encode(&baseRequestBody{
		SiteID:   c.SiteID,
		SitePass: c.SitePass,
		ShopID:   c.ShopID,
		ShopPass: c.ShopPass,
	}, baseForms)
	if err != nil {
		return nil, err
	}

	additinalForms := url.Values{}

	if err := shiftjis_transformer.EncodeToShiftJISFromUTF8(body); err != nil {
		return nil, err
	}

	err = parser.Encoder().Encode(body, additinalForms)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.APIHost, path),
		strings.NewReader(fmt.Sprintf("%s&%s", baseForms.Encode(), additinalForms.Encode())),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	curl, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		fmt.Println(curl.String())
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		fmt.Println("response.body", string(bodyBytes))
	}

	bodyBytes, err = shiftjis_transformer.DecodeToUTF8FromShiftJIS(bodyBytes)
	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)

	q, err := url.ParseQuery(bodyString)
	if err != nil {
		return nil, err
	}

	// NOTE:
	// レスポンスにURLが含まれている場合、&などのパラメータが抜け落ちてしまう可能性が高いので
	//　個別にparseして、URLを追加する
	splitedBody := strings.Split(bodyString, "RedirectUrl=")
	if len(splitedBody) > 1 {
		q.Set("RedirectUrl", splitedBody[1])
	}

	err = parser.Decoder().Decode(respBody, q)
	if err != nil {
		return nil, err
	}

	err = parser.ParseError(respBody)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) doOpenAPI(
	path string,
	jsonReq interface{},
	respBody interface{},
	errRespBody interface{},
) (*http.Response, error) {
	jsonReqBytes, err := json.Marshal(jsonReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json request: %w", err)
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", c.OpenAPIHost, path),
		bytes.NewBuffer(jsonReqBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")

	authotizationToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ShopID, c.ShopPass)))
	httpReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", authotizationToken))

	if c.Debug {
		curl, err := http2curl.GetCurlCommand(httpReq)
		if err != nil {
			return nil, fmt.Errorf("failed to get curl command: %w", err)
		}
		fmt.Println(curl.String())
	}

	httpRes, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}

	bodyBytes, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if httpRes.StatusCode >= 300 {
		err = json.Unmarshal(bodyBytes, errRespBody)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response body: %w", err)
		}
		return httpRes, nil
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}

	err = json.Unmarshal(bodyBytes, respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return httpRes, nil
}
