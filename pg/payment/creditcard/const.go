package creditcard

const (
	apiHostSandbox                = "https://pt01.mul-pay.jp"
	apiHostProduction             = "https://p01.mul-pay.jp"
	apiHostTest                   = "http://p01.mul-pay.jp"
	entryBankAccountPath          = "payment/BankAccountEntry.idPass"
	getResultEntryBankAccountPath = "payment/BankAccountTranResult.idPass"
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
