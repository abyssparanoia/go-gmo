package payment

import "net/http"

func newTestClient() *Client {
	cli, _ := NewClient("siteID", "sitePass", "shopID", "shopPass", false)
	cli.APIHost = apiHostTest
	cli.SetHTTPClient(http.DefaultClient)
	return cli
}
