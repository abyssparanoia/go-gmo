package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// PostpayEntryTranRequest ... postpay entry tran request
type PostpayEntryTranRequest struct {
	OrderID string `schema:"OrderID" validate:"required,max=27"`
	Amount  int    `schema:"Amount" validate:"required"`
	Tax     int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *PostpayEntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// PostpayEntryTranResponse ... postpay entry tran response
type PostpayEntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// PostpayEntryTran ... postpay entry tran
func (cli *Client) PostpayEntryTran(
	req *PostpayEntryTranRequest,
) (*PostpayEntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PostpayEntryTranResponse{}
	_, err := cli.do(postpayEntryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
