package payment

const (
	apiHostSandbox                = "https://pt01.mul-pay.jp"
	apiHostProduction             = "https://p01.mul-pay.jp"
	apiHostTest                   = "http://p01.mul-pay.jp"
	entryBankAccountPath          = "payment/BankAccountEntry.idPass"
	getResultEntryBankAccountPath = "payment/BankAccountTranResult.idPass"
	entryTranPath                 = "payment/EntryTran.idPass"
	execTranPath                  = "payment/ExecTran.idPass"
	alterTranPath                 = "payment/AlterTran.idPass"
	changeTranPath                = "payment/ChangeTran.idPass"
	saveMemberPath                = "payment/SaveMember.idPass"
	updateMemberPath              = "payment/UpdateMember.idPass"
	deleteMemberPath              = "payment/DeleteMember.idPass"
	searchMemberPath              = "payment/SearchMember.idPass"
	saveCardPath                  = "payment/SaveCard.idPass"
	deleteCardPath                = "payment/DeleteCard.idPass"
	searchCardPath                = "payment/SearchCard.idPass"
	entryTranGANBPath             = "payment/EntryTranGANB.idPass"
	execTranGANGPath              = "payment/ExecTranGANB.idPass"
	searchTradeMultiPath          = "payment/SearchTradeMulti.idPass"
	payEasyEntryTranPath          = "payment/EntryTranPayEasy.idPass"
	payEasyExecTranPath           = "payment/ExecTranPayEasy.idPass"
	convenienceStoreEntryTranPath = "payment/EntryTranCvs.idPass"
	convenienceStoreExecTranPath  = "payment/ExecTranCvs.idPass"
	convenienceStoreCancelPath    = "payment/CvsCancel.idPass"
	payPayEntryTranPath           = "payment/EntryTranPaypay.idPass"
)

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

// JobCD ... job cd type
type JobCD string

const (
	// JobCDCheck ... check
	JobCDCheck JobCD = "CHECK"
	// JobCDCapture ... capture
	JobCDCapture JobCD = "CAPTURE"
	// JobCDAuth ... auth
	JobCDAuth JobCD = "AUTH"
	// JobCDSauth ... sauth
	JobCDSauth JobCD = "SAUTH"
	// JobCDSales ... sales
	JobCDSales JobCD = "SALES"
	// JobCDCancel ... cancel
	JobCDCancel JobCD = "CANCEL"
)

func (j JobCD) String() string {
	return string(j)
}

// TradeMultiStatus ... search trade multi status
type TradeMultiStatus string

const (
	// TradeMultiStatusUnprocessed ... unprocessed
	TradeMultiStatusUnprocessed TradeMultiStatus = "UNPROCESSED"
	// TradeMultiStatusTrading ... proccessing
	TradeMultiStatusTrading TradeMultiStatus = "TRADING"
	// TradeMultiStatusPaysuccess ... success
	TradeMultiStatusPaysuccess TradeMultiStatus = "PAYSUCCESS"
	// TradeMultiStatusStop ... stop
	TradeMultiStatusStop TradeMultiStatus = "STOP"
	// TradeMultiStatusExpired ... expired
	TradeMultiStatusExpired TradeMultiStatus = "EXPIRED"
)

func (s TradeMultiStatus) String() string {
	return string(s)
}

// PayType ... pay type
type PayType string

const (
	// PayEasyPayType ... convenience store
	CvsPayType PayType = "3"
	// PayEasyPayType ... pay easy
	PayEasyPayType PayType = "4"
	// IDNetPayType ... iD net
	IDNetPayType PayType = "6"
	// GANBPayType ... GMO Aozora Net Bank
	GANBPayType PayType = "36"
)

// WebhookResultResponseStatus ... webhook result response status
type WebhookResultResponseStatus int

const (
	// WebhookResultResponseStatusOK ... ok
	WebhookResultResponseStatusOK WebhookResultResponseStatus = iota
	// WebhookResultResponseStatusFailed ... failed
	WebhookResultResponseStatusFailed
)

type WebhookResultPaymentSlipStatus string

const (
	// WebhookResultPaymentSlipStatusPaysuccess ... success
	WebhookResultPaymentSlipStatusPaysuccess WebhookResultPaymentSlipStatus = "PAYSUCCESS"
	// WebhookResultPaymentSlipStatusCancel ... cancel
	WebhookResultPaymentSlipStatusCancel WebhookResultPaymentSlipStatus = "CANCEL"
)

type ConvenienceStoreCode string

const (
	ConvenienceStoreCodeSevenEleven ConvenienceStoreCode = "00007"
	ConvenienceStoreCodeLawson      ConvenienceStoreCode = "10001"
	ConvenienceStoreCodeFamilyMart  ConvenienceStoreCode = "10002"
	ConvenienceStoreCodeMiniStop    ConvenienceStoreCode = "10005"
	ConvenienceStoreCodeSeikoMart   ConvenienceStoreCode = "10008"
)
