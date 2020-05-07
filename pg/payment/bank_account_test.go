package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"gopkg.in/go-playground/assert.v1"
)

func TestEntryBancAccount(t *testing.T) {

	expected := &EntryBankAccountResponse{
		TranID:   "tranID",
		Token:    "token",
		StartURL: "startURL",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := url.Values{}
		_ = parser.Encoder.Encode(expected, form)
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.Write([]byte(form.Encode()))
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

func TestGetResultEntryBankAccount(t *testing.T) {
	expected := &GetResultEntryBankAccountResponse{
		TranID:                "tranID",
		SiteID:                "siteID",
		MemberID:              "memberID",
		Status:                ResultEntryBankAccountStatusEntry,
		BankCode:              "0000",
		BranchCode:            "1234567",
		AccountType:           "1",
		AccountNumber:         "1234567",
		AccountName:           "YAMADA TAROU",
		AccountIdentification: "",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := url.Values{}
		_ = parser.Encoder.Encode(expected, form)
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.Write([]byte(form.Encode()))
	}))
	defer ts.Close()
	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

	cli, _ := NewClient("siteID", "sitePass", "shopID", "shopPass", false)
	cli.APIHost = apiHostTest
	req := &GetResultEntryBankAccountRequest{
		TranID: "tranID",
	}

	result, _ := cli.GetResultEntryBankAccount(req)

	assert.Equal(t, expected, result)
}
