package aozorabank

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type TransferDetails struct {
	TransferStatus string `json:"transferStatus"`
}

type TransferStatusRequest struct {
	AccountID     string `json:"accountId" validate:"required,min=12,max=29"`
	QueryKeyClass string `json:"queryKeyClass" validate:"required,len=1"`
	ApplyNo       string `json:"applyNo" validate:"omitempty,len=16"`
}

type TransferStatusResponse struct {
	AcceptanceKeyClass string             `json:"acceptanceKeyClass"`
	BaseDate           string             `json:"baseDate"`
	BaseTime           string             `json:"baseTime"`
	Count              string             `json:"count"`
	TransferDetails    []*TransferDetails `json:"transferDetails"`
}

func (r *TransferStatusRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) TransferStatus(
	req *TransferStatusRequest,
) (*TransferStatusResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &TransferStatusResponse{}
	if _, err := cli.doGet("transfer/status", reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}
