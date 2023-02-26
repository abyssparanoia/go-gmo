package aozorabank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/imdario/mergo"
)

// Client ... gmo pg remittance API client
type Client struct {
	HTTPClient  *http.Client
	APIHost     string
	accessToken string
}

// NewClient ... new client
func NewClient(
	sandBox bool,
	accessToken string,
) (*Client, error) {
	if accessToken == "" || len(accessToken) > 128 {
		return nil, fmt.Errorf("invalid access token, accessToken=%s", accessToken)
	}

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
		APIHost:     apiHost,
		accessToken: accessToken,
	}, nil
}

func (c *Client) doPost(
	header http.Header,
	path string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {
	return do(c.HTTPClient, c.accessToken, c.APIHost, header, path, http.MethodPost, body, respBody)
}

func (c *Client) doGet(
	path string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {
	values := url.Values{}
	for k, v := range body {
		values.Add(k, fmt.Sprintf("%s", v))
	}

	return do(c.HTTPClient, c.accessToken, c.APIHost, nil, fmt.Sprintf("%s?%s", path, values.Encode()), http.MethodGet, body, respBody)
}

func do(
	cli *http.Client,
	accessToken string,
	apiHost string,
	header http.Header,
	path string,
	method string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {

	fmt.Println("path: ", path)
	requestBodyMap := map[string]interface{}{}
	if err := mergo.Map(&requestBodyMap, &body); err != nil {
		return nil, err
	}

	requestBodyBytes, err := json.Marshal(requestBodyMap)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", apiHost, path),
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", accessToken)
	for k, values := range header {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	var resp *http.Response
	backoffCfg := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)
	err = backoff.Retry(func() (err error) {
		resp, err = cli.Do(req)
		if err != nil {
			return err
		}
		return nil
	}, backoffCfg)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if contains := bytes.Contains(bodyBytes, []byte("errorCode")); contains {
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
