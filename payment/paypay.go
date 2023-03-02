package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

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
