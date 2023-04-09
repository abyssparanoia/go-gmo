package aozorabank

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
	"net/http"
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
	reqMap, err := converter.StructToJsonTagMap(req)
	reqMap[grantTypeName] = grantTypeAuthorizationCode
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
	res := &CreateTokenResponse{}
	if _, err := cli.doPost(header, fmt.Sprintf("%s/token", authPathV1), reqMap, res); err != nil {
		return nil, err
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
