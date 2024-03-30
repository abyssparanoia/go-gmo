package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// EntryTranGANBRequest ... entry tran for bank transfer request type
type EntryTranGANBRequest struct {
	OrderID string `schema:"OrderID" validate:"required,max=27"`
	Amount  int    `schema:"Amount" validate:"required"`
	Tax     int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *EntryTranGANBRequest) Validate() error {
	return validate.Struct(r)
}

// EntryTranGANBResponse ... entry tran for bank transfer response type
type EntryTranGANBResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// EntryTranGANB ... entry tran for bank transfer
func (cli *Client) EntryTranGANB(
	req *EntryTranGANBRequest,
) (*EntryTranGANBResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &EntryTranGANBResponse{}
	_, err := cli.do(entryTranGANBPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ExecTranGANBRequest ... execc tran for bank transfer request type
type ExecTranGANBRequest struct {
	AccessID                  string `schema:"AccessID" validate:"required"`
	AccessPass                string `schema:"AccessPass" validate:"required"`
	OrderID                   string `schema:"OrderID" validate:"required"`
	ClientField1              string `schema:"ClientField1,omitempty"`
	ClientField2              string `schema:"ClientField2,omitempty"`
	ClientField3              string `schema:"ClientField3,omitempty"`
	AccountHolderOptionalName string `schema:"AccountHolderOptionalName,omitempty"`
	TradeDays                 int    `schema:"TradeDays,omitempty"`
	TradeReason               string `schema:"TradeReason,omitempty"`
	TradeClientName           string `schema:"TradeClientName,omitempty"`
	TradeClientMailaddress    string `schema:"TradeClientMailaddress,omitempty"`
}

// Validate ... validate
func (r *ExecTranGANBRequest) Validate() error {
	return validate.Struct(r)
}

// ExecTranGANBResponse ... exec tran for bank transfer response type
type ExecTranGANBResponse struct {
	AccessID          string `schema:"AccessID,omitempty"`
	BankCode          string `schema:"BankCode,omitempty"`
	BankName          string `schema:"BankName,omitempty"`
	BranchCode        string `schema:"BranchCode,omitempty"`
	BranchName        string `schema:"BranchName,omitempty"`
	AccountType       string `schema:"AccountType,omitempty"`
	AccountNumber     string `schema:"AccountNumber,omitempty"`
	AccountName       string `schema:"AccountName,omitempty"`
	AccountHolderName string `schema:"AccountHolderName,omitempty"`
	AvailableDate     string `schema:"AvailableDate,omitempty"`
	ErrCode           string `schema:"ErrCode,omitempty"`
	ErrInfo           string `schema:"ErrInfo,omitempty"`
}

// ExecTranGANB ... exec tran for bank transfer
func (cli *Client) ExecTranGANB(
	req *ExecTranGANBRequest,
) (*ExecTranGANBResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	res := &ExecTranGANBResponse{}
	_, err := cli.do(execTranGANGPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
