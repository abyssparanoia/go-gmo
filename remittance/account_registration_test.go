package remittance

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountRegistration(t *testing.T) {

	expected := &AccountRegistrationResponse{
		BankID: "bankID",
		Method: "1",
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
	cli.SetHTTPClient(http.DefaultClient)
	req := &AccountRegistrationRequest{
		Method:        BankAccountRegistrationMethodRegister,
		BankID:        "bankID",
		BankCode:      "0001",
		BranchCode:    "312",
		AccountType:   "1",
		AccountNumber: "1234567",
		AccountName:   "タナカタロウ",
	}
	result, err := cli.AccountRegistration(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}
