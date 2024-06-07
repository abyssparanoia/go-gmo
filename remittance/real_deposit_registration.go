package remittance

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type RealDepositRegistrationRequest struct {
	DepositID string `json:"Deposit_ID" validate:"required,max=27"`
	BankID    string `json:"Bank_ID" validate:"required,max=60"`
	Amount    string `json:"Amount" validate:"required,max=6"` // max 1,000,000
}

func (r *RealDepositRegistrationRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type RealDepositRegistrationResponse struct {
	DepositID string        `json:"Deposit_ID"`
	BankID    string        `json:"Bank_ID"`
	Amount    string        `json:"Amount"`
	Result    DepositResult `json:"Result"`
}

func (cli *Client) RealDepositRegistration(
	req *RealDepositRegistrationRequest,
) (*RealDepositRegistrationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &RealDepositRegistrationResponse{}
	_, err = cli.do(realDepositRegistrationPath, reqMap, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
