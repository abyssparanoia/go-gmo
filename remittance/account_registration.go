package remittance

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type AccountRegistrationRequest struct {
	Method              BankAccountRegistrationMethod `json:"Method" validate:"required"`
	BankID              string                        `json:"Bank_ID" validate:"required"`
	BankCode            string                        `json:"Bank_Code,omitempty"`
	BranchCode          string                        `json:"Branch_Code,omitempty"`
	AccountType         string                        `json:"Account_Type,omitempty"` // 1 - Normal, 2 - Current, 4 - Savings
	AccountNumber       string                        `json:"Account_Number,omitempty"`
	AccountName         string                        `json:"Account_Name,omitempty"`
	BranchCodeJPBank    string                        `json:"Branch_Code_Jpbank,omitempty"`
	AccountNumberJPBank string                        `json:"Account_Number_Jpbank,omitempty"`
	Free                string                        `json:"Free,omitempty"`
}

func (r *AccountRegistrationRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	if r.BankCode == "9900" {
		accountNumberJPbank := r.AccountNumber
		branchCodeJPbank := r.BranchCode
		r.AccountNumber = ""
		r.BranchCode = ""
		r.AccountNumberJPBank = accountNumberJPbank
		r.BranchCodeJPBank = branchCodeJPbank
	} else {
		r.BranchCodeJPBank = ""
		r.AccountNumberJPBank = ""
	}
	return nil
}

type AccountRegistrationResponse struct {
	BankID string                        `json:"Bank_ID"`
	Method BankAccountRegistrationMethod `json:"Method"`
}

func (cli *Client) AccountRegistration(
	req *AccountRegistrationRequest,
) (*AccountRegistrationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &AccountRegistrationResponse{}
	_, err = cli.do(accountRegistrationPath, reqMap, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
