package aozorabank

import (
	"github.com/bxcodec/faker"
	"net/http"
)

func fakeData[T any]() *T {
	ret := new(T)
	if err := faker.FakeData(ret); err != nil {
		panic(err)
	}
	return ret
}

func newTestAuthClient(clientID, clientSecret string, apiHost APIHostType) *AuthClient {
	cli, _ := NewAuthClient(clientID, clientSecret, apiHost)
	cli.SetHTTPClient(http.DefaultClient)
	return cli
}

func newTestClient(apiHost APIHostType) *Client {
	cli, _ := NewClient(apiHost)
	cli.SetHTTPClient(http.DefaultClient)
	return cli
}
