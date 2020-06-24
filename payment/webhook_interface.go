package payment

// WebhookResultEntryBankAccountRequest ... webhook request parameter
type WebhookResultEntryBankAccountRequest struct {
	TransactionID         string                       `schema:"TransactionID"`
	SiteID                string                       `schema:"SiteID"`
	MemberID              string                       `schema:"MemberID"`
	Result                ResultEntryBankAccountStatus `schema:"Result"`
	BankCode              string                       `schema:"BankCode"`
	BranchCode            string                       `schema:"BranchCode"`
	AccountType           string                       `schema:"AccountType"`
	AccountNumber         string                       `schema:"AccountNumber"`
	ErrCode               string                       `schema:"ErrCode"`
	ErrDetail             string                       `schema:"ErrDetail"`
	AccountIdentification string                       `schema:"AccountIdentification"`
}
