package aozorabank

import (
	"context"

	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

const (
	grantTypeName              = "grant_type"
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeRefreshToken      = "refresh_token"
)

type (
	CreateTokenRequest struct {
		Code             string           `json:"code" validate:"required,min=1,max=128"`
		RedirectURI      string           `json:"redirect_uri" validate:"required,uri,min=1,max=256"`
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

func (cli *AuthClient) CreateToken(
	ctx context.Context,
	req *CreateTokenRequest,
) (*CreateTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	reqMap[grantTypeName] = grantTypeAuthorizationCode
	res := &CreateTokenResponse{}
	if _, err := cli.doPost("token", req.ClientSecretType, reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type (
	RefreshTokenRequest struct {
		RefreshToken     string           `json:"refresh_token" validate:"required,min=1,max=128"`
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

func (cli *AuthClient) RefreshToken(
	ctx context.Context,
	req *RefreshTokenRequest,
) (*RefreshTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	reqMap[grantTypeName] = grantTypeRefreshToken
	res := &RefreshTokenResponse{}
	if _, err := cli.doPost("token", req.ClientSecretType, reqMap, res); err != nil {
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

func (cli *AuthClient) GetUserInfo(
	ctx context.Context,
	req *GetUserInfoRequest,
) (*GetUserInfoResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}

	res := &GetUserInfoResponse{}
	if _, err := cli.doGet("userinfo", req.AccessToken, reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}
