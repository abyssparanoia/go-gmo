package remittance

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type DepositRegistrationRequest struct {
	Method    DepositRegistrationMethod `json:"Method" validate:"required"`
	DepositID string                    `json:"Deposit_ID" validate:"required,max=27"`
	BankID    string                    `json:"Bank_ID"` // if method is 1, this field is required
	Amount    string                    `json:"Amount"`  // max 1,000,000. if method is 1, this field is required
}

func (r *DepositRegistrationRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type DepositRegistrationResponse struct {
	DepositID string                    `json:"Deposit_ID"`
	BankID    string                    `json:"Bank_ID"`
	Method    DepositRegistrationMethod `json:"Method"`
	Amount    string                    `json:"Amount"`
	BankFee   string                    `json:"Bank_Fee"`
}

func (cli *Client) DepositRegistration(
	req *DepositRegistrationRequest,
) (*DepositRegistrationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &DepositRegistrationResponse{}
	_, err = cli.do(depositRegistrationPath, reqMap, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
