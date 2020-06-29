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
