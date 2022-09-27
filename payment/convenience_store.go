package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// ConvenienceStoreEntryTranRequest ... convenience store entry tran request
type ConvenienceStoreEntryTranRequest struct {
	OrderID string `schema:"OrderID" validate:"required,max=27"`
	Amount  int    `schema:"Amount" validate:"required"`
	Tax     int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *ConvenienceStoreEntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// ConvenienceStoreEntryTranResponse ... convenience store entry tran response
type ConvenienceStoreEntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// ConvenienceStoreEntryTran ... convenience store entry tran
func (cli *Client) ConvenienceStoreEntryTran(
	req *ConvenienceStoreEntryTranRequest,
) (*ConvenienceStoreEntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &ConvenienceStoreEntryTranResponse{}
	_, err := cli.do(convenienceStoreEntryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ConvenienceStoreExecTranRequest ... convenience store exec tran request
type ConvenienceStoreExecTranRequest struct {
	AccessID        string               `schema:"AccessID" validate:"required"`
	AccessPass      string               `schema:"AccessPass" validate:"required"`
	OrderID         string               `schema:"OrderID" validate:"required"`
	Convenience     ConvenienceStoreCode `schema:"Convenience" validate:"required"`
	CustomerName    string               `schema:"CustomerName" validate:"required"`
	CustomerKana    string               `schema:"CustomerKana" validate:"required"`
	TelNo           string               `schema:"TelNo" validate:"required,max=13"`
	PaymentTermDay  int                  `schema:"PaymentTermDay,omitempty"`
	MailAddress     string               `schema:"MailAddress,omitempty"`
	ShopMailAddress string               `schema:"ShopMailAddress,omitempty"`
	ReserveNo       int                  `schema:"ReserveNo,omitempty"`
	MemberNo        int                  `schema:"MemberNo,omitempty"`
	RegisterDisp1   string               `schema:"RegisterDisp1,omitempty"`
	RegisterDisp2   string               `schema:"RegisterDisp2,omitempty"`
	RegisterDisp3   string               `schema:"RegisterDisp3,omitempty"`
	RegisterDisp4   string               `schema:"RegisterDisp4,omitempty"`
	RegisterDisp5   string               `schema:"RegisterDisp5,omitempty"`
	RegisterDisp6   string               `schema:"RegisterDisp6,omitempty"`
	RegisterDisp7   string               `schema:"RegisterDisp7,omitempty"`
	RegisterDisp8   string               `schema:"RegisterDisp8,omitempty"`
	ReceiptsDisp1   string               `schema:"ReceiptsDisp1,omitempty"`
	ReceiptsDisp2   string               `schema:"ReceiptsDisp2,omitempty"`
	ReceiptsDisp3   string               `schema:"ReceiptsDisp3,omitempty"`
	ReceiptsDisp4   string               `schema:"ReceiptsDisp4,omitempty"`
	ReceiptsDisp5   string               `schema:"ReceiptsDisp5,omitempty"`
	ReceiptsDisp6   string               `schema:"ReceiptsDisp6,omitempty"`
	ReceiptsDisp7   string               `schema:"ReceiptsDisp7,omitempty"`
	ReceiptsDisp8   string               `schema:"ReceiptsDisp8,omitempty"`
	ReceiptsDisp9   string               `schema:"ReceiptsDisp9,omitempty"`
	ReceiptsDisp10  string               `schema:"ReceiptsDisp10,omitempty"`
	ReceiptsDisp11  string               `schema:"ReceiptsDisp11,omitempty"`
	ReceiptsDisp12  string               `schema:"ReceiptsDisp12,omitempty"`
	ReceiptsDisp13  string               `schema:"ReceiptsDisp13,omitempty"`
	ClientField1    string               `schema:"ClientField1,omitempty"`
	ClientField2    string               `schema:"ClientField2,omitempty"`
	ClientField3    string               `schema:"ClientField3,omitempty"`
	ClientFieldFlag string               `schema:"ClientFieldFlag,omitempty"`
}

// Validate ... validate
func (r *ConvenienceStoreExecTranRequest) Validate() error {
	return validate.Struct(r)
}

// ConvenienceStoreExecTranResponse ... convenience store exec tran response
type ConvenienceStoreExecTranResponse struct {
	OrderID      string `schema:"OrderID,omitempty"`
	Convenience  int    `schema:"Convenience,omitempty"`
	ConfNo       string `schema:"ConfNo,omitempty"`
	ReceiptNo    string `schema:"ReceiptNo,omitempty"`
	PaymentTerm  int    `schema:"PaymentTerm,omitempty"`
	TranDate     int    `schema:"TranDate,omitempty"`
	ReceiptUrl   string `schema:"ReceiptUrl,omitempty"`
	CheckString  string `schema:"CheckString,omitempty"`
	ClientField1 string `schema:"ClientField1,omitempty"`
	ClientField2 string `schema:"ClientField2,omitempty"`
	ClientField3 string `schema:"ClientField3,omitempty"`
	ErrCode      string `schema:"ErrCode,omitempty"`
	ErrInfo      string `schema:"ErrInfo,omitempty"`
}

// ConvenienceStoreExecTran ... convenience store exec tran
func (cli *Client) ConvenienceStoreExecTran(
	req *ConvenienceStoreExecTranRequest,
) (*ConvenienceStoreExecTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &ConvenienceStoreExecTranResponse{}
	_, err := cli.do(convenienceStoreExecTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ConvenienceStoreCancelRequest ... convenience store cancel transaction request
type ConvenienceStoreCancelRequest struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
	OrderID    string `schema:"OrderID" validate:"required,max=27"`
}

// Validate ... validate
func (r *ConvenienceStoreCancelRequest) Validate() error {
	return validate.Struct(r)
}

// ConvenienceStoreCancelResponse ... convenience store cancel transaction response
type ConvenienceStoreCancelResponse struct {
	OrderID string `schema:"OrderID,omitempty"`
	Status  string `schema:"Status,omitempty"`
	ErrCode string `schema:"ErrCode,omitempty"`
	ErrInfo string `schema:"ErrInfo,omitempty"`
}

// ConvenienceStoreCancel ... convenience store cancel transaction
func (cli *Client) ConvenienceStoreCancel(
	req *ConvenienceStoreCancelRequest,
) (*ConvenienceStoreCancelResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &ConvenienceStoreCancelResponse{}
	_, err := cli.do(convenienceStoreCancelPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
