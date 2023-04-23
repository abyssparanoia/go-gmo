package aozorabank

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
)

// AuthClient ... gmo pg remittance API auth client
type AuthClient struct {
	cli          *http.Client
	clientID     string
	clientSecret string
	apiHost      string
}

const (
	authPathV1 = "auth/v1"
)

// NewAuthClient ... new auth client
func NewAuthClient(
	clientID string,
	clientSecret string,
	apiHostType ApiHostType,
) (*AuthClient, error) {
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("invalid client id or client secret, clientID=%s, clientSecret=%s", clientID, clientSecret)
	}

	var apiHost string
	switch apiHostType {
	case ApiHostTypeSandbox:
		apiHost = apiHostSandbox
	case ApiHostTypeStaging:
		apiHost = apiHostStaging
	case ApiHostTypeProduction:
		apiHost = apiHostProduction
	case ApiHostTypeTest:
		apiHost = apiHostTest
	default:
		return nil, fmt.Errorf("invalid api host type, apiHostType=%d", apiHostType)
	}

	return &AuthClient{
		cli: &http.Client{
			Timeout: time.Second * 30,
		},
		clientID:     clientID,
		clientSecret: clientSecret,
		apiHost:      fmt.Sprintf("%s/%s", apiHost, authPathV1),
	}, nil
}

func (c *AuthClient) doPost(
	path string,
	clientSecretType ClientSecretType,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {
	values := url.Values{}
	for k, v := range body {
		values.Add(k, fmt.Sprintf("%s", v))
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", c.apiHost, path),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to new request, err=%w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if clientSecretType == ClientSecretTypeBasic {
		auth := fmt.Sprintf("%s:%s", c.clientID, c.clientSecret)
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth))))
	}

	var resp *http.Response
	backoffCfg := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)
	err = backoff.Retry(func() (err error) {
		resp, err = c.cli.Do(req)
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

	if isError(resp.StatusCode) {
		errResp := &AuthErrorResponse{}
		if err := json.Unmarshal(bodyBytes, errResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response, bodyBytes=%s,  err=%w", string(bodyBytes), err)
		}
		return nil, errResp
	}

	if err := json.Unmarshal(bodyBytes, respBody); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, err=%w", err)
	}

	return resp, nil
}

func (c *AuthClient) doGet(
	path string,
	accessToken string,
	body map[string]interface{},
	respBody interface{},
) (*http.Response, error) {
	values := url.Values{}
	for k, v := range body {
		values.Add(k, fmt.Sprintf("%s", v))
	}

	requestBodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body, err=%w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s?%s", c.apiHost, path, values.Encode()),
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to new request, err=%w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	var resp *http.Response
	backoffCfg := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)
	err = backoff.Retry(func() (err error) {
		resp, err = c.cli.Do(req)
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

	if isError(resp.StatusCode) {
		errResp := &AuthErrorResponse{}
		if err := json.Unmarshal(bodyBytes, errResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response, bodyBytes=%s,  err=%w", string(bodyBytes), err)
		}
		return nil, errResp
	}

	if err := json.Unmarshal(bodyBytes, respBody); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, err=%w", err)
	}

	return resp, nil
}
