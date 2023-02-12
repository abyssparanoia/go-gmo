package aozorabank

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestTransferStatus(
	t *testing.T,
) {
	expected := &TransferStatusResponse{
		AcceptanceKeyClass: "1",
		BaseDate:           "2018-08-10",
		BaseTime:           "10:00:00+09:00",
		Count:              "500",
		TransferDetails:    nil,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respBody, _ := json.Marshal(expected)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)
	}))
	defer ts.Close()
	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(ts.URL)
	}
	defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

	cli, _ := NewClient(false, "testAccessToken")
	cli.APIHost = apiHostTest
	req := &TransferStatusRequest{
		AccountID:     "111111111111",
		QueryKeyClass: "1",
		ApplyNo:       "1111111111111111",
	}
	result, err := cli.TransferStatus(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}
