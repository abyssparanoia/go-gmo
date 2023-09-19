package aozorabank

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateToken(t *testing.T) {
	testcases := map[string]struct {
		request  *CreateTokenRequest
		host     string
		path     string
		header   http.Header
		expected *CreateTokenResponse
		wantErr  error
	}{
		"ok: client_secret_type=basic": {
			request: &CreateTokenRequest{
				Code:             "code",
				RedirectURI:      "https://example.com",
				ClientSecretType: ClientSecretTypeBasic,
			},
			path: "http://api.gmo-aozora.com/ganb/api/auth/v1/token",
			header: http.Header{
				"Authorization": []string{"Basic Y2xpZW50X2lkOmNsaWVudF9zZWNyZXQ="},
				"Content-Type":  []string{"application/x-www-form-urlencoded"},
			},
			expected: &CreateTokenResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				Scope:        "scope",
				TokenType:    "token_type",
				ExpiresIn:    3600,
			},
		},
		"ok: client_secret_type=post": {
			request: &CreateTokenRequest{
				Code:             "code",
				RedirectURI:      "https://example.com",
				ClientSecretType: ClientSecretTypePost,
			},
			path: "http://api.gmo-aozora.com/ganb/api/auth/v1/token",
			header: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			expected: &CreateTokenResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				Scope:        "scope",
				TokenType:    "token_type",
				ExpiresIn:    3600,
			},
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				assert.Equal(t, tc.path, req.URL.String())
				assert.Equal(t, tc.header, req.Header)

				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewAuthClient("client_id", "client_secret", APIHostTypeTest)
			result, err := cli.CreateToken(context.TODO(), tc.request)
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

func TestRefreshToken(t *testing.T) {
	testcases := map[string]struct {
		request  *RefreshTokenRequest
		host     string
		path     string
		header   http.Header
		expected *RefreshTokenResponse
		wantErr  error
	}{
		"ok: client_secret_type=basic": {
			request: &RefreshTokenRequest{
				RefreshToken:     "refresh_token",
				ClientSecretType: ClientSecretTypeBasic,
			},
			path: "http://api.gmo-aozora.com/ganb/api/auth/v1/token",
			header: http.Header{
				"Authorization": []string{"Basic Y2xpZW50X2lkOmNsaWVudF9zZWNyZXQ="},
				"Content-Type":  []string{"application/x-www-form-urlencoded"},
			},
			expected: &RefreshTokenResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				Scope:        "scope",
				TokenType:    "token_type",
				ExpiresIn:    3600,
			},
		},
		"ok: client_secret_type=post": {
			request: &RefreshTokenRequest{
				RefreshToken:     "refresh_token",
				ClientSecretType: ClientSecretTypePost,
			},
			path: "http://api.gmo-aozora.com/ganb/api/auth/v1/token",
			header: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			expected: &RefreshTokenResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				Scope:        "scope",
				TokenType:    "token_type",
				ExpiresIn:    3600,
			},
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				assert.Equal(t, tc.path, req.URL.String())
				assert.Equal(t, tc.header, req.Header)

				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewAuthClient("client_id", "client_secret", APIHostTypeTest)
			result, err := cli.RefreshToken(context.TODO(), tc.request)
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

func TestGetUserInfo(t *testing.T) {
	testcases := map[string]struct {
		request  *GetUserInfoRequest
		host     string
		path     string
		header   http.Header
		expected *GetUserInfoResponse
		wantErr  error
	}{
		"ok": {
			request: &GetUserInfoRequest{
				AccessToken: "access_token",
			},
			path: "http://api.gmo-aozora.com/ganb/api/auth/v1/userinfo?",

			header: http.Header{
				"Authorization": []string{"Bearer access_token"},
			},
			expected: &GetUserInfoResponse{
				Sub: "sub",
				Iss: "iss",
				Sup: "sup",
			},
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				assert.Equal(t, tc.path, req.URL.String())
				assert.Equal(t, tc.header, req.Header)

				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewAuthClient("client_id", "client_secret", APIHostTypeTest)
			result, err := cli.GetUserInfo(context.TODO(), tc.request)
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
