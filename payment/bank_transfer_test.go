package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"gopkg.in/go-playground/assert.v1"
)

func TestEntryGANBTran(t *testing.T) {

	expected := &EntryTranGANBResponse{
		AccessID:   "accessID",
		AccessPass: "accessPass",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := url.Values{}
		_ = parser.Encoder().Encode(expected, form)
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.Write([]byte(form.Encode()))
	}))
	defer ts.Close()
	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

	cli := newTestClient()
	req := &EntryTranGANBRequest{
		OrderID: "orderID",
		Amount:  4214144,
		Tax:     2414,
	}
	result, _ := cli.EntryTranGANB(req)
	assert.Equal(t, expected, result)
}

func TestExecGANBTran(t *testing.T) {

	expected := &ExecTranGANBResponse{
		AccessID: "accessID",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := url.Values{}
		_ = parser.Encoder().Encode(expected, form)
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.Write([]byte(form.Encode()))
	}))
	defer ts.Close()
	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

	cli := newTestClient()
	req := &ExecTranGANBRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
	}
	result, _ := cli.ExecTranGANB(req)
	assert.Equal(t, expected, result)
}
