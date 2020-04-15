package payment

// EntryResultBankAccountStatus ... entry bank account status
type EntryResultBankAccountStatus string

const (
	// EntryResultBankAccountStatusEntry ... entry complete
	EntryResultBankAccountStatusEntry EntryResultBankAccountStatus = "ENTRY"
	// EntryResultBankAccountStatusStart ... entry start
	EntryResultBankAccountStatusStart EntryResultBankAccountStatus = "START"
	// EntryResultBankAccountStatusTerm ... result confirm
	EntryResultBankAccountStatusTerm EntryResultBankAccountStatus = "TERM"
	// EntryResultBankAccountStatusSuccess ... entry success
	EntryResultBankAccountStatusSuccess EntryResultBankAccountStatus = "SUCCESS"
	// EntryResultBankAccountStatusFail ... entry fail because of bank reason
	EntryResultBankAccountStatusFail EntryResultBankAccountStatus = "FAIL"
	// EntryResultBankAccountStatusUnprocessed ... entry fail because of other reason
	EntryResultBankAccountStatusUnprocessed EntryResultBankAccountStatus = "UNPROCESSED"
)

func (s EntryResultBankAccountStatus) String() string {
	return string(s)
}

// WebhookEntryResultBankAccountRequest ... webhook request parameter
type WebhookEntryResultBankAccountRequest struct {
	TransactionID         string                       `json:"TransactionID"`
	SiteID                string                       `json:"SiteID"`
	MemberID              string                       `json:"MemberID"`
	Result                EntryResultBankAccountStatus `json:"Result"`
	BankCode              string                       `json:"BankCode"`
	BranchCode            string                       `json:"BranchCode"`
	AccountType           string                       `json:"AccountType"`
	ErrCode               string                       `json:"ErrCode"`
	ErrDetail             string                       `json:"ErrDetail"`
	AccountIdentification string                       `json:"AccountIdentification"`
}
