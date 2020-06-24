package deferred

import (
	"net/http"
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
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
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

func (c *Client) do() (*http.Response, error) {
	return nil, nil
}
