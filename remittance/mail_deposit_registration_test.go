package remittance

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMailDepositRegistration(t *testing.T) {

	expected := &MailDepositRegistrationResponse{
		DepositID: "depositID",
		Method:    "1",
		Amount:    "1000",
		Expire:    "20200101",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respBody, _ := json.Marshal(expected)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)
	}))
	defer ts.Close()
	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

	cli, _ := NewClient("shopID", "shopPass", false)
	cli.APIHost = apiHostTest
	req := &MailDepositRegistrationRequest{
		Method:                 "1",
		DepositID:              "depositID",
		Amount:                 1000,
		MailAddress:            "test@gogmo.co.jp",
		MailDepositAccountName: "TEST NAME",
		Expire:                 "30",
	}
	result, err := cli.MailDepositRegistration(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}
