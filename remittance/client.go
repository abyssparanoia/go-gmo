package remittance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
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
			Timeout: time.Second * 30,
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

	fmt.Println(string(requestBodyBytes))

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.APIHost, path),
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
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
