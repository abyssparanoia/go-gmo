package payment

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestEntryBancAccount(t *testing.T) {

	expected := &EntryBankAccountResponse{
		TrainID:  "trainID",
		Token:    "token",
		StartURL: "startURL",
		ErrCode:  "errCode",
		ErrInfo:  "errInfo",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		js, _ := json.Marshal(expected)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}))
	defer ts.Close()
	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

	cli, _ := NewClient("siteID", "sitePass", "shopID", "shopPass", false)
	cli.APIHost = apiHostTest
	req := &EntryBankAccountRequest{
		MemberID:         "memberID",
		MemberName:       "member name",
		CreateMember:     "1",
		RetURL:           "http://returl",
		BankCode:         "0000",
		BranchCode:       "1234567",
		AccountType:      "1",
		AccountNumber:    "1234567",
		AccountName:      "YAMADA TAROU",
		AccountNameKanji: "山田 太郎",
		ConsumerDevice:   "pc",
	}

	result, _ := cli.EntryBankAccount(req)

	assert.Equal(t, expected, result)
}
