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

// PayPaySalesRequest ... paypay sales request
type PayPaySalesRequest struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
	OrderID    string `schema:"OrderID" validate:"required"`
	Amount     int    `schema:"Amount" validate:"required"`
	Tax        int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *PayPaySalesRequest) Validate() error {
	return validate.Struct(r)
}

// PayPaySalesResponse ... paypay sales response
type PayPaySalesResponse struct {
	OrderID string           `schema:"OrderID,omitempty"`
	Amount  int              `schema:"Amount,omitempty"`
	Tax     int              `schema:"Tax,omitempty"`
	Status  TradeMultiStatus `schema:"Status,omitempty"`
	ErrCode string           `schema:"ErrCode,omitempty"`
	ErrInfo string           `schema:"ErrInfo,omitempty"`
}

// PayPaySales ... paypay sales
func (cli *Client) PayPaySales(
	req *PayPaySalesRequest,
) (*PayPaySalesResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PayPaySalesResponse{}
	_, err := cli.do(payPaySalesPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PayPayCancelReturnRequest ... paypay cancel return request
type PayPayCancelReturnRequest struct {
	AccessID     string `schema:"AccessID" validate:"required"`
	AccessPass   string `schema:"AccessPass" validate:"required"`
	OrderID      string `schema:"OrderID" validate:"required"`
	CancelAmount int    `schema:"CancelAmount" validate:"required"`
	CacnelTax    int    `schema:"CacnelTax,omitempty"`
}

// Validate ... validate
func (r *PayPayCancelReturnRequest) Validate() error {
	return validate.Struct(r)
}

// PayPaySalesResponse ... paypay sales response
type PayPayCancelReturnResponse struct {
	OrderID      string           `schema:"OrderID,omitempty"`
	Status       TradeMultiStatus `schema:"Status,omitempty"`
	CancelAmount int              `schema:"CancelAmount,omitempty"`
	CacnelTax    int              `schema:"CacnelTax,omitempty"`
	ErrCode      string           `schema:"ErrCode,omitempty"`
	ErrInfo      string           `schema:"ErrInfo,omitempty"`
}

// PayPaySales ... paypay sales
func (cli *Client) PayPayCancelReturn(
	req *PayPayCancelReturnRequest,
) (*PayPayCancelReturnResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PayPayCancelReturnResponse{}
	_, err := cli.do(payPaySalesPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
