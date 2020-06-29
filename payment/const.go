package payment

const (
	apiHostSandbox                = "https://pt01.mul-pay.jp"
	apiHostProduction             = "https://p01.mul-pay.jp"
	apiHostTest                   = "http://p01.mul-pay.jp"
	entryBankAccountPath          = "payment/BankAccountEntry.idPass"
	getResultEntryBankAccountPath = "payment/BankAccountTranResult.idPass"
	entryTranPath                 = "payment/EntryTran.idPass"
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
