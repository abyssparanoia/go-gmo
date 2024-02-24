package payment

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/abyssparanoia/go-gmo/internal/pkg/shiftjis_transformer"
	"github.com/cenkalti/backoff"
)

// Client ... gmo pg payment API client
type Client struct {
	HTTPClient *http.Client
	SiteID     string
	SitePass   string
	ShopID     string
	ShopPass   string
	APIHost    string
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

	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
		SiteID:   siteID,
		SitePass: sitePass,
		ShopID:   shopID,
		ShopPass: shopPass,
		APIHost:  apiHost,
	}, nil
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.HTTPClient = httpClient
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
	var resp *http.Response
	backoffCfg := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 4)
	err = backoff.Retry(func() (err error) {
		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}, backoffCfg)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyBytes, err = shiftjis_transformer.DecodeToUTF8FromShiftJIS(bodyBytes)
	if err != nil {
		return nil, err
	}

	q, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, err
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
