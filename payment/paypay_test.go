package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestPayPayEntryTran(t *testing.T) {

	expected := &PayPayEntryTranResponse{
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
	req := &PayPayEntryTranRequest{
		OrderID: "orderID",
		Amount:  4214144,
		JobCD:   JobCDAuth,
		Tax:     2414,
	}
	result, _ := cli.PayPayEntryTran(req)
	assert.Equal(t, expected, result)
}

func TestPayPayExecTran(t *testing.T) {

	expected := &PayPayExecTranResponse{
		OrderID: "orderID",
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

	req := &PayPayExecTranRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
		RetURL:     "http://localhost:8000/paypay/callback",
	}
	result, _ := cli.PayPayExecTran(req)
	assert.Equal(t, expected, result)
}

func TestPayPaySales(t *testing.T) {

	expected := &PayPaySalesResponse{
		OrderID: "orderID",
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

	req := &PayPaySalesRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
		Amount:     1000,
	}
	result, _ := cli.PayPaySales(req)
	assert.Equal(t, expected, result)
}

func TestPayPayCancelReturn(t *testing.T) {

	expected := &PayPayCancelReturnResponse{
		OrderID: "orderID",
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

	req := &PayPayCancelReturnRequest{
		AccessID:     "accessID",
		AccessPass:   "accessPass",
		OrderID:      "orderID",
		CancelAmount: 1000,
	}
	result, _ := cli.PayPayCancelReturn(req)
	assert.Equal(t, expected, result)
}
