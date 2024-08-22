package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestPayEasyEntryTran(t *testing.T) {

	expected := &PayEasyEntryTranResponse{
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
	req := &PayEasyEntryTranRequest{
		OrderID: "orderID",
		Amount:  4214144,
		Tax:     2414,
	}
	result, _ := cli.PayEasyEntryTran(req)
	assert.Equal(t, expected, result)
}

func TestPayEasyExecTran(t *testing.T) {

	expected := &PayEasyExecTranResponse{
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

	req := &PayEasyExecTranRequest{
		AccessID:     "accessID",
		AccessPass:   "accessPass",
		OrderID:      "orderID",
		CustomerName: "田中太郎",
		CustomerKana: "タナカタロウ",
		TelNo:        "012345678910",
	}
	result, _ := cli.PayEasyExecTran(req)
	assert.Equal(t, expected, result)
}
