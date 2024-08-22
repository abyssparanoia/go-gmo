package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestClient_ConvenienceStoreEntryTran(t *testing.T) {
	expected := &ConvenienceStoreEntryTranResponse{
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
	req := &ConvenienceStoreEntryTranRequest{
		OrderID: "orderID",
		Amount:  4214144,
		Tax:     2414,
	}
	result, _ := cli.ConvenienceStoreEntryTran(req)
	assert.Equal(t, expected, result)
}

func TestClient_ConvenienceStoreExecTran(t *testing.T) {

	expected := &ConvenienceStoreExecTranResponse{
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

	req := &ConvenienceStoreExecTranRequest{
		AccessID:     "accessID",
		AccessPass:   "accessPass",
		OrderID:      "orderID",
		Convenience:  ConvenienceStoreCodeSevenEleven,
		CustomerName: "田中太郎",
		CustomerKana: "タナカタロウ",
		TelNo:        "012345678910",
	}
	result, _ := cli.ConvenienceStoreExecTran(req)
	assert.Equal(t, expected, result)
}

func TestClient_ConvenienceStoreCancel(t *testing.T) {
	expected := &ConvenienceStoreCancelResponse{
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
	req := &ConvenienceStoreCancelRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
	}
	result, _ := cli.ConvenienceStoreCancel(req)
	assert.Equal(t, expected, result)
}
