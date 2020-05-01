package creditcard

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// EntryTranRequest ... entry tran request type
type EntryTranRequest struct {
	OrderID  string `json:"OrderID" validate:"required,max=27"`
	JobCD    JobCD  `json:"JobCd" validate:"required"`
	ItemCode string `json:"ItemCode,omitempty"`
	Amount   int    `json:"Amount"`
	Tax      int    `json:"Tax,omitempty"`
}

// Validate ... validate
func (r *EntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// EntryTranResponse ... entry tran response
type EntryTranResponse struct {
	AccessID   string `json:"AccessID"`
	AccessPass string `json:"AccessPass"`
	ErrCode    string `json:"ErrCode"`
	ErrInfo    string `json:"ErrInfo"`
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
