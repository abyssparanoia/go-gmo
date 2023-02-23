package aozorabank

const (
	apiHostSandbox    = "https://api.sunabar.gmo-aozora.com/corporation/v1"
	apiHostProduction = "https://api.gmo-aozora.com/ganb/api/corporation/v1"
	apiHostTest       = "http://api.gmo-aozora.com/ganb/api/corporation/v1"
)

const (
	IdempotencyKeyHeaderKey = "Idempotency-Key"
)

type QueryKeyClass int

const (
	QueryKeyClassTransferApplies   QueryKeyClass = 1
	QueryKeyClassTransferQueryBulk QueryKeyClass = 2
)

type RequestTransferStatus int

const (
	RequestTransferStatusApplying                RequestTransferStatus = 2
	RequestTransferStatusReturned                RequestTransferStatus = 3
	RequestTransferStatusDismiss                 RequestTransferStatus = 4
	RequestTransferStatusExpired                 RequestTransferStatus = 5
	RequestTransferStatusApprovalCancelled       RequestTransferStatus = 8
	RequestTransferStatusInReserve               RequestTransferStatus = 11
	RequestTransferStatusInProgress              RequestTransferStatus = 12
	RequestTransferStatusRetrying                RequestTransferStatus = 13
	RequestTransferStatusCompleted               RequestTransferStatus = 20
	RequestTransferStatusFundsReturned           RequestTransferStatus = 22
	RequestTransferStatusTransferReturning       RequestTransferStatus = 24
	RequestTransferStatusTransferReturnCompleted RequestTransferStatus = 25
	RequestTransferStatusTransferReturnFailed    RequestTransferStatus = 26
	RequestTransferStatusFailed                  RequestTransferStatus = 40
)

type RequestTransferClass int

const (
	RequestTransferClassAll                  RequestTransferClass = 1
	RequestTransferClassTransferApplyingOnly RequestTransferClass = 2
	RequestTransferClassTransferAcceptsOnly  RequestTransferClass = 3
)

type RequestTransferTerm int

const (
	RequestTransferTermTransferApplyingApplyDatetime RequestTransferTerm = 1
	RequestTransferTermTransferDesignatedDate        RequestTransferTerm = 2
)

type TransferDateHolidayCode int

const (
	TransferDateHolidayCodeNextBusinessDay  = 1
	TransferDateHolidayCodePreviousBusiness = 2
	TransferDateHolidayCodeErrorReturn      = 3
)

type AccountTypeCode int

const (
	AccountTypeCodeOrdinary AccountTypeCode = 1
	AccountTypeCodeChecking AccountTypeCode = 2
	AccountTypeCodeSavings  AccountTypeCode = 4
	AccountTypeCodeOther    AccountTypeCode = 9
)

type ResultCode int

const (
	ResultCodeCompleted  ResultCode = 1
	ResultCodeIncomplete ResultCode = 2
	ResultCodeExpired    ResultCode = 8
)
