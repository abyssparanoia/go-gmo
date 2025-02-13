package payment

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestEntryTran(t *testing.T) {

	expected := &EntryTranResponse{
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

func TestExecTran(t *testing.T) {

	expected := &ExecTranResponse{
		ACS:    "acs",
		ACSUrl: "acs_url",
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
	req := &ExecTranRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
		MemberID:   "memberID",
		SeqMode:    "1",
		CardSeq:    1,
	}
	result, _ := cli.ExecTran(req)
	assert.Equal(t, expected, result)
}

func TestExecTranWithToken(t *testing.T) {

	expected := &ExecTranResponse{
		ACS:    "acs",
		ACSUrl: "acs_url",
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
	req := &ExecTranRequestWithToken{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		OrderID:    "orderID",
		TokenType:  "1",
		Token:      "token",
	}
	result, _ := cli.ExecTranWithToken(req)
	assert.Equal(t, expected, result)
}

func TestAlterTran(t *testing.T) {

	expected := &AlterTranResponse{
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
	req := &AlterTranRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		JobCD:      JobCDSales,
		Amount:     1000,
	}
	result, _ := cli.AlterTran(req)
	assert.Equal(t, expected, result)
}

func TestChangeAlterTran(t *testing.T) {

	expected := &ChangeTranResponse{
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
	req := &ChangeTranRequest{
		AccessID:   "accessID",
		AccessPass: "accessPass",
		JobCD:      JobCDSales,
		Amount:     1000,
	}
	result, _ := cli.ChangeTran(req)
	assert.Equal(t, expected, result)
}

func TestSecureTran2(t *testing.T) {

	expected := &SecureTran2Response{
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
	req := &SecureTran2Request{
		AccessID:   "accessID",
		AccessPass: "accessPass",
	}
	result, _ := cli.SecureTran2(req)
	assert.Equal(t, expected, result)
}
