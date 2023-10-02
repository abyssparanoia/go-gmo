package aozorabank

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AuthClient_doPost(t *testing.T) {
	clientID := "client_id"
	clientSecret := "client_secret"
	client, _ := NewAuthClient(
		clientID,
		clientSecret,
		APIHostTypeTest,
	)

	testCases := map[string]struct {
		respBody interface{}
		expected *GetUserInfoResponse
		respErr  error
		wantErr  error
	}{
		"success": {
			respBody: &GetUserInfoResponse{
				Sub: "sub",
				Iss: "iss",
				Sup: "sup",
			},
			expected: &GetUserInfoResponse{
				Sub: "sub",
				Iss: "iss",
				Sup: "sup",
			},
		},
		"ng": {
			respErr: &AuthErrorResponse{ErrorCode: "code", ErrorDescription: "description", ErrorURI: "uri"},
			wantErr: &AuthErrorResponse{ErrorCode: "code", ErrorDescription: "description", ErrorURI: "uri"},
		},
		"ng: failed to unmarshal error response": {
			respErr: &ErrorResponse{ErrorCode: "code", ErrorMessage: "message"},
			wantErr: fmt.Errorf("failed to unmarshal error response, bodyBytes={\"errorCode\":\"code\",\"errorMessage\":\"message\",\"errorDetails\":null,\"transferErrorDetails\":null}\n"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tc.wantErr != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(tc.respErr)
					return
				}
				json.NewEncoder(w).Encode(tc.respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			actual := &GetUserInfoResponse{}
			_, err := client.doPost(ts.URL, ClientSecretTypeBasic, nil, actual)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_AuthClient_doGet(t *testing.T) {
	clientID := "client_id"
	clientSecret := "client_secret"
	accessToken := "access_token"
	client, _ := NewAuthClient(
		clientID,
		clientSecret,
		APIHostTypeTest,
	)

	testCases := map[string]struct {
		respBody interface{}
		expected *GetUserInfoResponse
		respErr  error
		wantErr  error
	}{
		"success": {
			respBody: &GetUserInfoResponse{
				Sub: "sub",
				Iss: "iss",
				Sup: "sup",
			},
			expected: &GetUserInfoResponse{
				Sub: "sub",
				Iss: "iss",
				Sup: "sup",
			},
		},
		"ng": {
			respErr: &AuthErrorResponse{ErrorCode: "code", ErrorDescription: "description", ErrorURI: "uri"},
			wantErr: &AuthErrorResponse{ErrorCode: "code", ErrorDescription: "description", ErrorURI: "uri"},
		},
		"ng: failed to unmarshal error response": {
			respErr: &ErrorResponse{ErrorCode: "code", ErrorMessage: "message"},
			wantErr: fmt.Errorf("failed to unmarshal error response, bodyBytes={\"errorCode\":\"code\",\"errorMessage\":\"message\",\"errorDetails\":null,\"transferErrorDetails\":null}\n"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tc.wantErr != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(tc.respErr)
					return
				}
				json.NewEncoder(w).Encode(tc.respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			actual := &GetUserInfoResponse{}
			_, err := client.doGet(ts.URL, accessToken, nil, actual)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

func Test_unmarshalError(t *testing.T) {

	clientID := "client_id"
	clientSecret := "client_secret"
	client, _ := NewAuthClient(
		clientID,
		clientSecret,
		APIHostTypeTest,
	)

	testCases := map[string]struct {
		errBytes []byte
		want     error
	}{
		"ok: ErrorResponse": {
			errBytes: []byte(`{"errorCode":"WG_ERR_300","errorMessage":"under maintenance"}`),
			want: &ErrorResponse{
				ErrorCode:    "WG_ERR_300",
				ErrorMessage: "under maintenance",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := client.unmarshalError(tc.errBytes)
			assert.Equal(t, tc.want.Error(), got.Error())
		})
	}
}
