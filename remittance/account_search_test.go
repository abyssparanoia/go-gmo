package remittance

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountSearch(t *testing.T) {

	expected := &AccountSearchResponse{
		BankID:        "bankID",
		DeleteFlag:    BankAccountDeleteFlagActive,
		BankName:      "bankName",
		BankCode:      "0001",
		BranchName:    "branchName",
		BranchCode:    "312",
		AccountType:   "1",
		AccountNumber: "1234567",
		AccountName:   "タナカタロウ",
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
	req := &AccountSearchRequest{
		BankID: "bankID",
	}
	result, err := cli.AccountSearch(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}
