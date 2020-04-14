package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/validate"
)

// EntryBankAccountRequest ... input parameter for Entrying bank account
type EntryBankAccountRequest struct {
	MemberID   string `json:"MemberID" validate:"required"`
	MemberName string `json:"MemberName,omitempty"`
	// 0 ... do not create member, 1 ... create member
	CreateMember string `json:"CreateMember" validate:"required,oneof=0 1"`
	RetURL       string `json:"RetURL" validate:"required,url"`
	BankCode     string `json:"BankCode" validate:"required,len=4"`
	BranchCode   string `json:"BranchCode,omitempty"`
	// 1 ... normal type, 2 ... touza type
	AccountType      string `json:"AccountType,omitempty"`
	AccountNumber    string `json:"AccountNumber,omitempty"`
	AccountName      string `json:"AccountName,omitempty"`
	AccountNameKanji string `json:"AccountNameKanji,omitempty"`
	ConsumerDevice   string `json:"ConsumerDevice" validate:"required,oneof=i ez sb pc"`
}

// EntryBankAccountResponse ... response parameter
type EntryBankAccountResponse struct {
	TrainID  string `json:"TrainID"`
	Token    string `json:"token"`
	StartURL string `json:"StartUrl"`
	ErrCode  string `json:"ErrCode"`
	ErrInfo  string `json:"ErrInfo"`
}

// EntryBankAccount ... Entry bank account
func (cli *Client) EntryBankAccount(
	req *EntryBankAccountRequest,
) (res *EntryBankAccountResponse, err error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	_, err = cli.do(entryBankAccountPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
