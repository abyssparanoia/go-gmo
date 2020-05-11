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
	ClientField1              string `schema:"ClientField1"`
	ClientField2              string `schema:"ClientField2"`
	ClientField3              string `schema:"ClientField3"`
	AccountHolderOptionalName string `schema:"AccountHolderOptionalName"`
	TradeDays                 int    `schema:"TradeDays"`
	TradeReason               string `schema:"TradeReason"`
	TradeClientName           string `schema:"TradeClientName"`
	TradeClientMailaddress    string `schema:"TradeClientMailaddress"`
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

// SearchTradeMultiRequest ... search trade multi requst
type SearchTradeMultiRequest struct {
	OrderID string `schema:"OrderID" validate:"required"`
	PayType string `schema:"PayType" validate:"required"`
}

// Validate ... validate
func (r *SearchTradeMultiRequest) Validate() error {
	return validate.Struct(r)
}

// SearchTradeMultiResponse ... search trade multi response type
type SearchTradeMultiResponse struct {
	Status                   TradeMultiStatus `schema:"Status,omitempty"`
	ProcessDate              string           `schema:"ProcessDate,omitempty"`
	AccessID                 string           `schema:"AccessID,omitempty"`
	AccessPass               string           `schema:"AccessPass,omitempty"`
	Amount                   int              `schema:"Amount,omitempty"`
	Tax                      int              `schema:"Tax,omitempty"`
	ClientField1             string           `schema:"ClientField1,omitempty"`
	ClientField2             string           `schema:"ClientField2,omitempty"`
	ClientField3             string           `schema:"ClientField3,omitempty"`
	PayType                  string           `schema:"PayType,omitempty"`
	GanbBankCode             string           `schema:"GanbBankCode,omitempty"`
	GanbBankName             string           `schema:"GanbBankName,omitempty"`
	GanbBranchCode           string           `schema:"GanbBranchCode,omitempty"`
	GanbBranchName           string           `schema:"GanbBranchName,omitempty"`
	GanbAccountType          string           `schema:"GanbAccountType,omitempty"`
	GanbAccountNumber        string           `schema:"GanbAccountNumber,omitempty"`
	GanbAccountHolderName    string           `schema:"GanbAccountHolderName,omitempty"`
	GanbExpireDays           int              `schema:"GanbExpireDays,omitempty"`
	GanbExpireDate           string           `schema:"GanbExpireDate,omitempty"`
	GanbTradeReason          string           `schema:"GanbTradeReason,omitempty"`
	GanbTradeClientName      string           `schema:"GanbTradeClientName,omitempty"`
	GanbTotalTransferAmount  int              `schema:"GanbTotalTransferAmount,omitempty"`
	GanbTotalTransferCount   int              `schema:"GanbTotalTransferCount,omitempty"`
	GanbLatestTransferAmount int              `schema:"GanbLatestTransferCount,omitempty"`
}

// SearchTradeMulti ... search trade multi
func (cli *Client) SearchTradeMulti(
	req *SearchTradeMultiRequest,
) (*SearchTradeMultiResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	res := &SearchTradeMultiResponse{}
	_, err := cli.do(searchTradeMultiPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
