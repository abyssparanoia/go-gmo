package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// PayEasyEntryTranRequest ... pay easy entry tran request
type PayEasyEntryTranRequest struct {
	OrderID string `schema:"OrderID" validate:"required,max=27"`
	Amount  int    `schema:"Amount" validate:"required"`
	Tax     int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *PayEasyEntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// PayEasyEntryTranResponse ... pay easy entry tran response
type PayEasyEntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// PayEasyEntryTran ... pay easy entry tran
func (cli *Client) PayEasyEntryTran(
	req *PayEasyEntryTranRequest,
) (*PayEasyEntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PayEasyEntryTranResponse{}
	_, err := cli.do(payEasyEntryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PayEasyExecTranRequest ... pay easy exec tran request
type PayEasyExecTranRequest struct {
	AccessID        string `schema:"AccessID" validate:"required"`
	AccessPass      string `schema:"AccessPass" validate:"required"`
	OrderID         string `schema:"OrderID" validate:"required"`
	CustomerName    string `schema:"CustomerName" validate:"required"`
	CustomerKana    string `schema:"CustomerKana" validate:"required"`
	TelNo           string `schema:"TelNo" validate:"required,max=13"`
	PaymentTermDay  int    `schema:"PaymentTermDay,omitempty"`
	MailAddress     string `schema:"MailAddress,omitempty"`
	ShopMailAddress string `schema:"ShopMailAddress,omitempty"`
	RegisterDisp1   string `schema:"RegisterDisp1,omitempty"`
	RegisterDisp2   string `schema:"RegisterDisp2,omitempty"`
	RegisterDisp3   string `schema:"RegisterDisp3,omitempty"`
	RegisterDisp4   string `schema:"RegisterDisp4,omitempty"`
	RegisterDisp5   string `schema:"RegisterDisp5,omitempty"`
	RegisterDisp6   string `schema:"RegisterDisp6,omitempty"`
	RegisterDisp7   string `schema:"RegisterDisp7,omitempty"`
	RegisterDisp8   string `schema:"RegisterDisp8,omitempty"`
	ReceiptsDisp1   string `schema:"ReceiptsDisp1,omitempty"`
	ReceiptsDisp2   string `schema:"ReceiptsDisp2,omitempty"`
	ReceiptsDisp3   string `schema:"ReceiptsDisp3,omitempty"`
	ReceiptsDisp4   string `schema:"ReceiptsDisp4,omitempty"`
	ReceiptsDisp5   string `schema:"ReceiptsDisp5,omitempty"`
	ReceiptsDisp6   string `schema:"ReceiptsDisp6,omitempty"`
	ReceiptsDisp7   string `schema:"ReceiptsDisp7,omitempty"`
	ReceiptsDisp8   string `schema:"ReceiptsDisp8,omitempty"`
	ReceiptsDisp9   string `schema:"ReceiptsDisp9,omitempty"`
	ReceiptsDisp10  string `schema:"ReceiptsDisp10,omitempty"`
	ReceiptsDisp11  string `schema:"ReceiptsDisp11,omitempty"`
	ReceiptsDisp12  string `schema:"ReceiptsDisp12,omitempty"`
	ReceiptsDisp13  string `schema:"ReceiptsDisp13,omitempty"`
	ClientField1    string `schema:"ClientField1,omitempty"`
	ClientField2    string `schema:"ClientField2,omitempty"`
	ClientField3    string `schema:"ClientField3,omitempty"`
	ClientFieldFlag string `schema:"ClientFieldFlag,omitempty"`
	PaymentType     string `schema:"PaymentType,omitempty"`
}

// Validate ... validate
func (r *PayEasyExecTranRequest) Validate() error {
	return validate.Struct(r)
}

// PayEasyExecTranResponse ... pay easy exec tran response
type PayEasyExecTranResponse struct {
	OrderID          string `schema:"OrderID,omitempty"`
	CustID           string `schema:"CustID,omitempty"`
	BkCode           string `schema:"BkCode,omitempty"`
	ConfNo           string `schema:"ConfNo,omitempty"`
	EncryptReceiptNo string `schema:"EncryptReceiptNo,omitempty"`
	PaymentTerm      string `schema:"PaymentTerm,omitempty"`
	TranDate         string `schema:"TranDate,omitempty"`
	CheckString      string `schema:"CheckString,omitempty"`
	ClientField1     string `schema:"ClientField1,omitempty"`
	ClientField2     string `schema:"ClientField2,omitempty"`
	ClientField3     string `schema:"ClientField3,omitempty"`
	PaymentURL       string `schema:"PaymentURL,omitempty"`
	ErrCode          string `schema:"ErrCode,omitempty"`
	ErrInfo          string `schema:"ErrInfo,omitempty"`
}

// PayEasyExecTran ... pay easy exec tran
func (cli *Client) PayEasyExecTran(
	req *PayEasyExecTranRequest,
) (*PayEasyExecTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PayEasyExecTranResponse{}
	_, err := cli.do(payEasyExecTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
