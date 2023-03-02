package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// PayPayEntryTranRequest ... paypay entry tran request
type PayPayEntryTranRequest struct {
	OrderID string `schema:"OrderID" validate:"required,max=27"`
	Amount  int    `schema:"Amount" validate:"required"`
	JobCD   JobCD  `schema:"JobCd" validate:"required"`
	Tax     int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *PayPayEntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// PayPayEntryTranResponse ... paypay entry tran response
type PayPayEntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// PayPayEntryTran ... paypay entry tran
func (cli *Client) PayPayEntryTran(
	req *PayPayEntryTranRequest,
) (*PayPayEntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PayPayEntryTranResponse{}
	_, err := cli.do(payPayEntryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PayPayExecTranRequest ... paypay exec tran request
type PayPayExecTranRequest struct {
	AccessID       string               `schema:"AccessID" validate:"required"`
	AccessPass     string               `schema:"AccessPass" validate:"required"`
	OrderID        string               `schema:"OrderID" validate:"required"`
	RetURL         string               `schema:"RetURL" validate:"required,max=256"`
	PaymentTermSec int64                `schema:"PaymentTermSec,omitempty"`
	TransitionType PayPayTransitionType `schema:"TransitionType,omitempty"`
	ClientField1   string               `schema:"ClientField1,omitempty"`
	ClientField2   string               `schema:"ClientField2,omitempty"`
	ClientField3   string               `schema:"ClientField3,omitempty"`
}

// Validate ... validate
func (r *PayPayExecTranRequest) Validate() error {
	return validate.Struct(r)
}

// PayPayExecTranResponse ... paypay exec tran response
type PayPayExecTranResponse struct {
	AccessID       string `schema:"AccessID,omitempty"`
	OrderID        string `schema:"OrderID,omitempty"`
	Token          string `schema:"Token,omitempty"`
	StartURL       string `schema:"StartURL,omitempty"`
	StartLimitDate string `schema:"StartLimitDate,omitempty"`
	ErrCode        string `schema:"ErrCode,omitempty"`
	ErrInfo        string `schema:"ErrInfo,omitempty"`
}

// PayPayExecTran ... paypay exec tran
func (cli *Client) PayPayExecTran(
	req *PayPayExecTranRequest,
) (*PayPayExecTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PayPayExecTranResponse{}
	_, err := cli.do(payPayExecTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
