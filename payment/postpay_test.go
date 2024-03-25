package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"gopkg.in/go-playground/assert.v1"
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

	cli, _ := NewClient("siteID", "sitePass", "shopID", "shopPass", false)
	cli.APIHost = apiHostTest
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

	cli, _ := NewClient("siteID", "sitePass", "shopID", "shopPass", false)
	cli.APIHost = apiHostTest
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
