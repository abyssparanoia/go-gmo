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

// WebhookResultExecTranGANBRequest ... webhook request parameter
type WebhookResultExecTranGANBRequest struct {
	ShopID                     string           `shema:"ShopID"`
	ShopPass                   string           `shema:"ShopPass"`
	AccessID                   string           `schema:"AccessID"`
	AccessPass                 string           `schema:"AccessPass"`
	OrderID                    string           `schema:"OrderID"`
	Status                     TradeMultiStatus `schema:"Status"`
	Amount                     int              `schema:"Amount"`
	Tax                        int              `schema:"Tax"`
	TranDate                   string           `schema:"TranDate"`
	ErrCode                    string           `schema:"ErrCode"`
	ErrDetail                  string           `schema:"ErrDetail"`
	PayType                    PayType          `schema:"PayType"`
	GANBProcessType            string           `schema:"GanbProcessType"`
	GANBRequestAmount          string           `schema:"GanbRequestAmount"`
	GANBExpireDate             string           `schema:"GanbExpireDate"`
	GANBTradeReason            string           `schema:"GanbTradeReason"`
	GANBTradeClientName        string           `schema:"GanbTradeClientName"`
	GANBTradeClientMailAddress string           `schema:"GanbTradeClientMailAddress"`
	GANBBankCode               string           `schema:"GanbBankCode"`
	GANBBankName               string           `schema:"GanbBankName"`
	GANBBranchCode             string           `schema:"GanbBranchCode"`
	GANBBranchName             string           `schema:"GanbBranchName"`
	GANBAccountType            string           `schema:"GanbAccountType"`
	GANBAccountNumber          string           `schema:"GanbAccountNumber"`
	GANBAccountHolderName      string           `schema:"GanbAccountHolderName"`
	GANBSettlementDate         string           `schema:"GanbSettlementDate"`
	GANBInAmount               string           `schema:"GanbInAmount"`
	GANBInClientName           string           `schema:"GanbInClientName"`
	GANBInRemittingBranchName  string           `schema:"GanbInRemittingBranchName"`
	GANBTotalTransferAmount    string           `schema:"GanbTotalTransferAmount"`
	GANBTotalTransferCount     string           `schema:"GanbTotalTransferCount"`
}

type WebhookResultConvenienceStoreRequest struct {
	ShopID       string               `shema:"ShopID"`
	ShopPass     string               `shema:"ShopPass"`
	AccessID     string               `schema:"AccessID"`
	AccessPass   string               `schema:"AccessPass"`
	OrderID      string               `schema:"OrderID"`
	Status       TradeMultiStatus     `schema:"Status"`
	JobCD        JobCD                `schema:"JobCd"`
	Amount       int                  `schema:"Amount"`
	Tax          int                  `schema:"Tax"`
	Currency     string               `schema:"Currency"`
	TranID       string               `schema:"TranID"`
	TranDate     string               `schema:"TranDate"`
	CvsCode      ConvenienceStoreCode `schema:"CvsCode"`
	CvsConfNo    string               `schema:"CvsConfNo"`
	CvsReceiptNo string               `schema:"CvsReceiptNo"`
	PaymentTerm  string               `schema:"PaymentTerm"`
	FinishDate   string               `schema:"FinishDate"`
	ReceiptDate  string               `schema:"ReceiptDate"`
	ErrCode      string               `schema:"ErrCode"`
	ErrInfo      string               `schema:"ErrInfo"`
	PayType      PayType              `schema:"PayType"`
}

// WebhookResultPaymentSlipRequest ... webhook request parameter
type WebhookResultPaymentSlipRequest struct {
	ShopID      string                         `shema:"ShopID"`
	ShopPass    string                         `shema:"ShopPass"`
	AccessID    string                         `schema:"AccessID"`
	AccessPass  string                         `schema:"AccessPass"`
	OrderID     string                         `schema:"OrderID"`
	Status      WebhookResultPaymentSlipStatus `schema:"Status"`
	JobCD       JobCD                          `schema:"JobCd"`
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
