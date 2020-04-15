package payment

// WebhookEntryResultBankAccountRequest ... webhook request parameter
type WebhookEntryResultBankAccountRequest struct {
	TransactionID         string `json:"TransactionID"`
	SiteID                string `json:"SiteID"`
	MemberID              string `json:"MemberID"`
	Result                string `json:"Result"`
	BankCode              string `json:"BankCode"`
	BranchCode            string `json:"BranchCode"`
	AccountType           string `json:"AccountType"`
	ErrCode               string `json:"ErrCode"`
	ErrDetail             string `json:"ErrDetail"`
	AccountIdentification string `json:"AccountIdentification"`
}
