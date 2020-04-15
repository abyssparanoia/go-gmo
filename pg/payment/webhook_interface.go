package payment

// ResultEntryBankAccountStatus ... entry bank account status
type ResultEntryBankAccountStatus string

const (
	// ResultEntryBankAccountStatusEntry ... entry complete
	ResultEntryBankAccountStatusEntry ResultEntryBankAccountStatus = "ENTRY"
	// ResultEntryBankAccountStatusStart ... entry start
	ResultEntryBankAccountStatusStart ResultEntryBankAccountStatus = "START"
	// ResultEntryBankAccountStatusTerm ... result confirm
	ResultEntryBankAccountStatusTerm ResultEntryBankAccountStatus = "TERM"
	// ResultEntryBankAccountStatusSuccess ... entry success
	ResultEntryBankAccountStatusSuccess ResultEntryBankAccountStatus = "SUCCESS"
	// ResultEntryBankAccountStatusFail ... entry fail because of bank reason
	ResultEntryBankAccountStatusFail ResultEntryBankAccountStatus = "FAIL"
	// ResultEntryBankAccountStatusUnprocessed ... entry fail because of other reason
	ResultEntryBankAccountStatusUnprocessed ResultEntryBankAccountStatus = "UNPROCESSED"
)

func (s ResultEntryBankAccountStatus) String() string {
	return string(s)
}

// WebhookResultEntryBankAccountRequest ... webhook request parameter
type WebhookResultEntryBankAccountRequest struct {
	TransactionID         string                       `json:"TransactionID"`
	SiteID                string                       `json:"SiteID"`
	MemberID              string                       `json:"MemberID"`
	Result                ResultEntryBankAccountStatus `json:"Result"`
	BankCode              string                       `json:"BankCode"`
	BranchCode            string                       `json:"BranchCode"`
	AccountType           string                       `json:"AccountType"`
	AccountNumber         string                       `json:"AccountNumber"`
	ErrCode               string                       `json:"ErrCode"`
	ErrDetail             string                       `json:"ErrDetail"`
	AccountIdentification string                       `json:"AccountIdentification"`
}
