package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// EntryTranRequest ... entry tran request type
type EntryTranRequest struct {
	OrderID  string `schema:"OrderID" validate:"required,max=27"`
	JobCD    JobCD  `schema:"JobCd" validate:"required"`
	ItemCode string `schema:"ItemCode,omitempty"`
	Amount   int    `schema:"Amount,omitempty"`
	Tax      int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *EntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// EntryTranResponse ... entry tran response
type EntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// EntryTran ... entry tran
func (cli *Client) EntryTran(
	req *EntryTranRequest,
) (*EntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &EntryTranResponse{}
	_, err := cli.do(entryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
