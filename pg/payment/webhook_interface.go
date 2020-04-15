package payment

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
