package aozorabank

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

const (
	authPathV1                 = "auth/v1"
	grantTypeName              = "grant_type"
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeRefreshToken      = "refresh_token"
)

type (
	AuthorizationRequest struct {
		ClientID     string `json:"client_id" validate:"required,min=1,max=128"`
		RedirectURI  string `json:"redirect_uri" validate:"required,uri,min=1,max=256"`
		ResponseType string `json:"response_type" validate:"required"`
		Scope        string `json:"scope" validate:"required,min=1,max=256"`
		State        string `json:"state" validate:"required,min=1,max=128"`
	}

	AuthorizationResponse struct{}
)

func (r *AuthorizationRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) Authorization(
	ctx context.Context,
	req *AuthorizationRequest,
) (*AuthorizationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &AuthorizationResponse{}
	if _, err := cli.doGet(http.Header{}, fmt.Sprintf("%s/authorization", authPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type (
	CreateTokenRequest struct {
		Code             string           `json:"code" validate:"required,min=1,max=128"`
		RedirectURI      string           `json:"redirect_uri" validate:"required,uri,min=1,max=256"`
		ClientID         string           `json:"client_id" validate:"min=1,max=128"`
		ClientSecret     string           `json:"client_secret" validate:"min=1,max=128"`
		ClientSecretType ClientSecretType `json:"-" validate:"required"`
	}
	CreateTokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
	}
	CreateTokenErrResponse struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorURI         string `json:"error_uri"`
	}
)

func (r *CreateTokenRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) CreateToken(
	ctx context.Context,
	req *CreateTokenRequest,
) (*CreateTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("client_id", req.ClientID)
	data.Set("client_secret", req.ClientSecret)
	data.Set("grant_type", grantTypeAuthorizationCode)
	data.Set("code", req.Code)
	data.Set("redirect_uri", req.RedirectURI)

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", cli.apiHost, fmt.Sprintf("%s/token", authPathV1)),
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to new request, err=%w", err)
	}

	header := http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	}
	request.Header = header

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if isError(response.StatusCode) {
		errResp := &CreateTokenErrResponse{}
		if err := json.Unmarshal(bodyBytes, errResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response, bodyBytes=%s,  err=%w", string(bodyBytes), err)
		}
		return nil, fmt.Errorf("failed to create token, err=%s, description=%s, uri=%s",
			errResp.Error,
			errResp.ErrorDescription,
			errResp.ErrorURI,
		)
	}

	res := &CreateTokenResponse{}
	if err := json.Unmarshal(bodyBytes, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, err=%w", err)
	}
	return res, nil
}

type (
	RefreshTokenRequest struct {
		RefreshToken     string           `json:"refresh_token" validate:"required,min=1,max=128"`
		ClientID         string           `json:"client_id" validate:"min=1,max=128"`
		ClientSecret     string           `json:"client_secret" validate:"min=1,max=128"`
		ClientSecretType ClientSecretType `json:"-" validate:"required"`
	}
	RefreshTokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
	}
)

func (r *RefreshTokenRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) RefreshToken(
	ctx context.Context,
	req *RefreshTokenRequest,
) (*RefreshTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	reqMap[grantTypeName] = grantTypeRefreshToken
	if err != nil {
		return nil, err
	}
	header := http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	}
	if req.ClientSecretType == ClientSecretTypeBasic {
		if req.ClientID == "" || req.ClientSecret == "" {
			return nil, fmt.Errorf("invalid client id or client secret, clientID=%s, clientSecret=%s", req.ClientID, req.ClientSecret)
		}
		auth := fmt.Sprintf("%s:%s", req.ClientID, req.ClientSecret)
		encoded := base64.StdEncoding.EncodeToString([]byte(auth))
		header.Add("Authorization", encoded)
	}
	res := &RefreshTokenResponse{}
	if _, err := cli.doPost(header, fmt.Sprintf("%s/token", authPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type (
	GetUserInfoRequest struct {
		AccessToken string `json:"-" validate:"required,min=1,max=128"`
	}
	GetUserInfoResponse struct {
		Sub string `json:"sub"`
		Iss string `json:"iss"`
		Sup string `json:"sup"`
	}
)

func (r *GetUserInfoRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) GetUserInfo(
	ctx context.Context,
	req *GetUserInfoRequest,
) (*GetUserInfoResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	header := http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", req.AccessToken)},
	}
	res := &GetUserInfoResponse{}
	if _, err := cli.doGet(header, fmt.Sprintf("%s/userinfo", authPathV1), nil, res); err != nil {
		return nil, err
	}
	return res, nil
}
