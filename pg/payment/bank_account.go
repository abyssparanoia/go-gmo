package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// EntryBankAccountRequest ... input parameter for Entrying bank account
type EntryBankAccountRequest struct {
	MemberID   string `schema:"MemberID" validate:"required"`
	MemberName string `schema:"MemberName,omitempty"`
	// 0 ... do not create member, 1 ... create member
	CreateMember string `schema:"CreateMember" validate:"required,oneof=0 1"`
	RetURL       string `schema:"RetURL" validate:"required,url"`
	BankCode     string `schema:"BankCode" validate:"required,len=4"`
	BranchCode   string `schema:"BranchCode,omitempty"`
	// 1 ... normal type, 2 ... touza type
	AccountType      string `schema:"AccountType,omitempty"`
	AccountNumber    string `schema:"AccountNumber,omitempty"`
	AccountName      string `schema:"AccountName,omitempty"`
	AccountNameKanji string `schema:"AccountNameKanji,omitempty"`
	ConsumerDevice   string `schema:"ConsumerDevice" validate:"required,oneof=i ez sb pc"`
}

// EntryBankAccountResponse ... response parameter
type EntryBankAccountResponse struct {
	TranID   string `schema:"TranID"`
	Token    string `schema:"Token"`
	StartURL string `schema:"StartUrl"`
	ErrCode  string `schema:"ErrCode"`
	ErrInfo  string `schema:"ErrInfo"`
}

// EntryBankAccount ... Entry bank account
func (cli *Client) EntryBankAccount(
	req *EntryBankAccountRequest,
) (*EntryBankAccountResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	res := &EntryBankAccountResponse{}
	_, err := cli.do(entryBankAccountPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetResultEntryBankAccountRequest ... input parameter for getting result entry bank account
type GetResultEntryBankAccountRequest struct {
	TranID string `schema:"TranID" validate:"required"`
}

// GetResultEntryBankAccountResponse ... response parameter
type GetResultEntryBankAccountResponse struct {
	TranID                string                       `schema:"TranID"`
	SiteID                string                       `schema:"SiteID"`
	MemberID              string                       `schema:"MemberID"`
	Status                ResultEntryBankAccountStatus `schema:"Status"`
	BankCode              string                       `schema:"BankCode"`
	BranchCode            string                       `schema:"BranchCode"`
	AccountType           string                       `schema:"AccountType"`
	AccountNumber         string                       `schema:"AccountNumber"`
	AccountName           string                       `schema:"AccountName"`
	ErrCode               string                       `schema:"ErrCode"`
	ErrDetail             string                       `schema:"ErrDetail"`
	AccountIdentification string                       `schema:"AccountIdentification"`
}

// GetResultEntryBankAccount ... get result of entry bank account
func (cli *Client) GetResultEntryBankAccount(
	req *GetResultEntryBankAccountRequest,
) (*GetResultEntryBankAccountResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	res := &GetResultEntryBankAccountResponse{}
	_, err := cli.do(getResultEntryBankAccountPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
