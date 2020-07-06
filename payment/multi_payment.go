package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// SearchTradeMultiRequest ... search trade multi requst
type SearchTradeMultiRequest struct {
	OrderID string  `schema:"OrderID" validate:"required"`
	PayType PayType `schema:"PayType" validate:"required"`
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
	ErrCode                  string           `schema:"ErrCode,omitempty"`
	ErrInfo                  string           `schema:"ErrInfo,omitempty"`
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
