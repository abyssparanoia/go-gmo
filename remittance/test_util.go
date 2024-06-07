package remittance

import "net/http"

func newTestClient() *Client {
	cli, _ := NewClient("shopID", "shopPass", false)
	cli.APIHost = apiHostTest
	cli.SetHTTPClient(http.DefaultClient)
	return cli
}
