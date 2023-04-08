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
	cli         *http.Client
	apiHost     string
	accessToken string
}

// NewClient ... new client
func NewClient(
	apiHostType ApiHostType,
) (*Client, error) {
	var apiHost string
	switch apiHostType {
	case ApiHostTypeSandbox:
		apiHost = apiHostSandbox
	case ApiHostTypeStaging:
		apiHost = apiHostStaging
	case ApiHostTypeProduction:
		apiHost = apiHostProduction
	default:
		return nil, fmt.Errorf("invalid api host type, apiHostType=%d", apiHostType)
	}

	return &Client{
		cli: &http.Client{
			Timeout: time.Second * 30,
		},
		apiHost: apiHost,
	}, nil
}

func (c *Client) doPost(
	header http.Header,
	path string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {
	return do(c.cli, c.accessToken, c.apiHost, header, path, http.MethodPost, body, respBody)
}

func (c *Client) doGet(
	header http.Header,
	path string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {
	values := url.Values{}
	for k, v := range body {
		values.Add(k, fmt.Sprintf("%s", v))
	}

	return do(c.cli, c.accessToken, c.apiHost, header, fmt.Sprintf("%s?%s", path, values.Encode()), http.MethodGet, body, respBody)
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

	req.Header = header
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
