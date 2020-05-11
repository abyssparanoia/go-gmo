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
