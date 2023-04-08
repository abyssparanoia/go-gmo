package aozorabank

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAuthorization(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		request  *AuthorizationRequest
		rawQuery string
		expected *AuthorizationResponse
		wantErr  error
	}{
		"ok": {
			request: &AuthorizationRequest{
				ClientID:     "111111111111",
				RedirectURI:  "https://example.com",
				ResponseType: "code",
				Scope:        "openid",
				State:        "state",
			},
			rawQuery: "client_id=111111111111&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&scope=openid&state=state",
			expected: &AuthorizationResponse{},
		},
		"ng: RedirectURI is not uri format": {
			request: &AuthorizationRequest{
				ClientID:     "111111111111",
				RedirectURI:  "dummy",
				ResponseType: "code",
				Scope:        "openid",
				State:        "state",
			},
			wantErr: fmt.Errorf("Key: 'AuthorizationRequest.RedirectURI' Error:Field validation for 'RedirectURI' failed on the 'uri' tag"),
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(tc.expected)
				assert.Equal(t, tc.rawQuery, r.URL.RawQuery)
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
			result, err := cli.Authorization(context.TODO(), tc.request)
			if tc.wantErr != nil {
				s := err.Error()
				fmt.Println(s)
				assert.Equal(t, tc.wantErr.Error(), err.Error())
				assert.Equal(t, tc.expected, nil)
			} else {
				assert.Equal(t, nil, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
