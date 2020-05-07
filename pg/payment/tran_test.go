package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"gopkg.in/go-playground/assert.v1"
)

func TestEntryTran(t *testing.T) {

	expected := &EntryTranResponse{
		AccessID:   "accessID",
		AccessPass: "accessPass",
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
	req := &EntryTranRequest{
		OrderID:  "orderID",
		JobCD:    JobCDAuth,
		ItemCode: "itemCode",
		Amount:   4214144,
		Tax:      2414,
	}
	result, _ := cli.EntryTran(req)
	assert.Equal(t, expected, result)
}
