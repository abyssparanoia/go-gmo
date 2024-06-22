package remittance

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type DepositSearchRequest struct {
	DepositID string `json:"Deposit_ID" validate:"required,max=27"`
}

func (r *DepositSearchRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type DepositSearchResponse struct {
	DepositID string                     `json:"Deposit_ID"`
	Amount    string                     `json:"Amount"`
	Free      string                     `json:"Free"`
	SelectKey string                     `json:"Select_Key"`
	Bank      *DepositSearchBankResponse `json:"bank"`
}

type DepositSearchBankResponse struct {
	BankID              string                  `json:"Bank_ID"`
	BankName            string                  `json:"Bank_Name"`
	BankCode            string                  `json:"Bank_Code"`
	BranchName          string                  `json:"Branch_Name"`
	BranchCode          string                  `json:"Branch_Code"`
	AccountType         string                  `json:"Account_Type"`
	AccountNumber       string                  `json:"Account_Number"`
	AccountName         string                  `json:"Account_Name"`
	BankFee             string                  `json:"Bank_Fee"`
	Result              DepositBankStatus       `json:"Result"`
	BranchCodeJpbank    string                  `json:"Branch_Code_Jpbank"`
	AccountNumberJpbank string                  `json:"Account_Number_Jpbank"`
	DepositDate         string                  `json:"Deposit_Date"`
	ResultDetail        DepositBankResultDetail `json:"Result_Detail"`
	ClientName          string                  `json:"Client_Name"`
}

func (cli *Client) DepositSearch(
	req *DepositSearchRequest,
) (*DepositSearchResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &DepositSearchResponse{}
	_, err = cli.do(depositSearchPath, reqMap, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
