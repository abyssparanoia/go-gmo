package payment

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSaveCard(t *testing.T) {

	expected := &SaveCardResponse{
		CardSeq: "0001",
		ErrCode: "errCode",
		ErrInfo: "errInfo",
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
	req := &SaveCardRequest{
		MemberID:     "memberID",
		CardNo:       "3131414141414141",
		Expire:       "0125",
		SecurityCode: "0000",
	}
	result, _ := cli.SaveCard(req)
	assert.Equal(t, expected, result)
}
