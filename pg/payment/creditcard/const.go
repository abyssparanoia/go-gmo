package creditcard

const (
	apiHostSandbox    = "https://pt01.mul-pay.jp"
	apiHostProduction = "https://p01.mul-pay.jp"
	apiHostTest       = "http://p01.mul-pay.jp"
	entryTranPath     = "payment/EntryTran.idPass"
)

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
