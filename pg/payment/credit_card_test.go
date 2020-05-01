package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"gopkg.in/go-playground/assert.v1"
)

func TestSaveCard(t *testing.T) {

	expected := &SaveCardResponse{
		CardSeq: "0001",
		ErrCode: "errCode",
		ErrInfo: "errInfo",
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
	req := &SaveCardRequest{
		MemberID:     "memberID",
		SeqMode:      "1",
		CardNo:       "3131414141414141",
		Expire:       "0125",
		SecurityCode: "0000",
	}
	result, _ := cli.SaveCard(req)
	assert.Equal(t, expected, result)
}

func TestDeleteCard(t *testing.T) {

	expected := &DeleteCardResponse{
		CardSeq: "0001",
		ErrCode: "errCode",
		ErrInfo: "errInfo",
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
	req := &DeleteCardRequest{
		MemberID: "memberID",
		CardSeq:  "0001",
	}
	result, _ := cli.DeleteCard(req)
	assert.Equal(t, expected, result)
}

func TestSearchCard(t *testing.T) {

	expected := &SearchCardResponse{
		Cards: []*SearchCardResponseDetail{
			{
				CardSeq: "0001",
			},
			{
				CardSeq: "0002",
			},
		},
		ErrCode: "errCode",
		ErrInfo: "errInfo",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := url.Values{}
		_ = parser.Encoder.Encode(&SearchCardResponseDetail{
			CardSeq: "0001|0002",
			ErrCode: "errCode",
			ErrInfo: "errInfo",
		}, form)
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
	req := &SearchCardRequest{
		MemberID: "memberID",
	}
	result, _ := cli.SearchCard(req)
	assert.Equal(t, expected, result)
}
