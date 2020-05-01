package creditcard

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSaveMember(t *testing.T) {

	expected := &SaveMemberResponse{
		MemberID: "memberID",
		ErrCode:  "errCode",
		ErrInfo:  "errInfo",
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

	cli, _ := NewClient("shopID", "shopPass", false)
	cli.APIHost = apiHostTest
	req := &SaveMemberRequest{
		MemberID:   "memberID",
		MemberName: "memberName",
	}
	result, _ := cli.SaveMember(req)
	assert.Equal(t, expected, result)
}
