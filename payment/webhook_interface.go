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

// WebhookResultPaymentSlipRequest ... webhook request parameter
type WebhookResultPaymentSlipRequest struct {
	ShopID      string                         `shema:"ShopID"`
	ShopPass    string                         `shema:"ShopPass"`
	AccessID    string                         `schema:"AccessID"`
	AccessPass  string                         `schema:"AccessPass"`
	OrderID     string                         `schema:"OrderID"`
	Status      WebhookResultPaymentSlipStatus `schema:"Status"`
	JobCd       JobCD                          `schema:"JobCd"`
	Amount      int                            `schema:"Amount"`
	Tax         int                            `schema:"Tax"`
	Currency    string                         `schema:"Currency"`
	Forward     string                         `schema:"Forward"`
	Method      int                            `schema:"Method"`
	PayTimes    int                            `schema:"PayTimes"`
	TranID      string                         `schema:"TranID"`
	Approve     int                            `schema:"Approve"`
	TranDate    string                         `schema:"TranDate"`
	PaymentTerm int                            `schema:"PaymentTerm"`
	ErrCode     string                         `schema:"ErrCode"`
	ErrDetail   string                         `schema:"ErrDetail"`
	PayType     PayType                        `schema:"PayType"`
}

// WebhookResultResponse ... webhook response parameter
type WebhookResultResponse struct {
	Status WebhookResultResponseStatus
}
