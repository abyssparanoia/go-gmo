package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// EntryTranRequest ... entry tran request type
type EntryTranRequest struct {
	OrderID  string `schema:"OrderID" validate:"required,max=27"`
	JobCD    JobCD  `schema:"JobCd" validate:"required"`
	ItemCode string `schema:"ItemCode,omitempty"`
	Amount   int    `schema:"Amount"`
	Tax      int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *EntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// EntryTranResponse ... entry tran response
type EntryTranResponse struct {
	AccessID   string `schema:"AccessID"`
	AccessPass string `schema:"AccessPass"`
	ErrCode    string `schema:"ErrCode"`
	ErrInfo    string `schema:"ErrInfo"`
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
