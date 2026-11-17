package remittance

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type AccountSearchRequest struct {
	BankID string `json:"Bank_ID" validate:"required,max=60"`
}

func (r *AccountSearchRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type AccountSearchResponse struct {
	BankID              string                `json:"Bank_ID"`
	DeleteFlag          BankAccountDeleteFlag `json:"Delete_Flag"`
	BankName            string                `json:"Bank_Name"`
	BankCode            string                `json:"Bank_Code"`
	BranchName          string                `json:"Branch_Name"`
	BranchCode          string                `json:"Branch_Code"`
	AccountType         string                `json:"Account_Type"` // 1 - Normal, 2 - Current, 4 - Savings
	AccountNumber       string                `json:"Account_Number"`
	AccountName         string                `json:"Account_Name"`
	Free                string                `json:"Free"`
	BranchCodeJPBank    string                `json:"Branch_Code_Jpbank"`
	AccountNumberJPBank string                `json:"Account_Number_Jpbank"`
}

func (cli *Client) AccountSearch(
	req *AccountSearchRequest,
) (*AccountSearchResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &AccountSearchResponse{}
	_, err = cli.do(accountSearchPath, reqMap, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
