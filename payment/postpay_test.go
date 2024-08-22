package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestPostpayEntryTran(t *testing.T) {

	expected := &PostpayEntryTranResponse{
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
	req := &PostpayEntryTranRequest{
		OrderID: "orderID",
		Amount:  4214144,
		Tax:     2414,
	}
	result, _ := cli.PostpayEntryTran(req)
	assert.Equal(t, expected, result)
}

func TestPostpayExecTran(t *testing.T) {

	expected := &PostpayExecTranResponse{
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
	req := &PostpayExecTranRequest{
		AccessID:            "accessID",
		AccessPass:          "accessPass",
		OrderID:             "orderID",
		CustomerOrderDate:   "20201224",
		CustomerName:        "決済太郎",
		CustomerNameKana:    "ケッサイタロウ",
		CustomerZipCode:     "1000000",
		CustomerAddress:     "東京都",
		CustomerTel1:        "09011111111",
		CustomerEmail1:      "sample1@test.com",
		CustomerPaymentType: PostpayCustomerPaymentTypeInvoiceSentSeparately,
		CustomerID:          "testid",
		DetailName:          "商品名",
		DetailPrice:         1000,
		DetailQuantity:      1,
		DetailBrand:         "ブランド",
		DetailCategory:      "カテゴリ",
	}
	result, _ := cli.PostpayExecTran(req)
	assert.Equal(t, expected, result)
}

func TestPostpayShippedTran(t *testing.T) {

	expected := &PostpayShippedTranResponse{
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
	req := &PostpayShippedTranRequest{
		AccessID:      "accessID",
		AccessPass:    "accessPass",
		OrderID:       "orderID",
		PDCompanyCode: PostpayPDCompanyCodeYamato,
		SlipNo:        "1234567890",
	}
	result, _ := cli.PostpayShippedTran(req)
	assert.Equal(t, expected, result)
}

func TestPostpayCancelTran(t *testing.T) {

	expected := &PostpayCancelTranResponse{
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
	req := &PostpayCancelTranRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
	}
	result, _ := cli.PostpayCancelTran(req)
	assert.Equal(t, expected, result)
}
