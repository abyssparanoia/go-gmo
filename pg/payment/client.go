package payment

import (
	"errors"
	"net/http"
)

// Client ... gmo pg payment API client
type Client struct {
	client   *http.Client
	SiteID   string
	SitePass string
	ShopID   string
	ShopPass string
	APIHost  string
}

// NewClient ... new client
func NewClient(
	siteID,
	sitePass,
	shopID,
	shopPass string,
	sandBox bool) (*Client, error) {
	if !(siteID != "" && sitePass != "" && shopID != "" && shopPass != "") {
		return nil, errors.New("Invalid parameters")
	}

	var apiHost string
	if sandBox {
		apiHost = apiHostSandbox
	} else {
		apiHost = apiHostProduction
	}

	return &Client{
		SiteID:   siteID,
		SitePass: sitePass,
		ShopID:   shopID,
		ShopPass: shopPass,
		APIHost:  apiHost,
	}, nil
}
