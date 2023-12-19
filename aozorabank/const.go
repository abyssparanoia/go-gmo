package aozorabank

const (
	apiHostSandbox    = "https://api.sunabar.gmo-aozora.com"
	apiHostStaging    = "https://stg-api.gmo-aozora.com/ganb/api"
	apiHostProduction = "https://api.gmo-aozora.com/ganb/api"
	apiHostTest       = "http://api.gmo-aozora.com/ganb/api"
)

type APIHostType int

const (
	APIHostTypeSandbox    APIHostType = 1
	APIHostTypeStaging    APIHostType = 2
	APIHostTypeProduction APIHostType = 3
	APIHostTypeTest       APIHostType = 4
)

const (
	IdempotencyKeyHeaderKey = "Idempotency-Key"
)

type QueryKeyClass int

const (
	QueryKeyClassTransferApplies   QueryKeyClass = 1
	QueryKeyClassTransferQueryBulk QueryKeyClass = 2
)

type TransferStatus int

const (
	TransferStatusApplying                TransferStatus = 2
	TransferStatusReturned                TransferStatus = 3
	TransferStatusDismiss                 TransferStatus = 4
	TransferStatusExpired                 TransferStatus = 5
	TransferStatusApprovalCancelled       TransferStatus = 8
	TransferStatusInReserve               TransferStatus = 11
	TransferStatusInProgress              TransferStatus = 12
	TransferStatusRetrying                TransferStatus = 13
	TransferStatusCompleted               TransferStatus = 20
	TransferStatusFundsReturned           TransferStatus = 22
	TransferStatusTransferReturning       TransferStatus = 24
	TransferStatusTransferReturnCompleted TransferStatus = 25
	TransferStatusTransferReturnFailed    TransferStatus = 26
	TransferStatusFailed                  TransferStatus = 40
)

type BulkTransferStatus int

const (
	BulkTransferStatusApplying             BulkTransferStatus = 2
	BulkTransferStatusReturned             BulkTransferStatus = 3
	BulkTransferStatusDismiss              BulkTransferStatus = 4
	BulkTransferStatusExpired              BulkTransferStatus = 5
	BulkTransferStatusApprovalCancelled    BulkTransferStatus = 8
	BulkTransferStatusInReserve            BulkTransferStatus = 11
	BulkTransferStatusInProgress           BulkTransferStatus = 12
	BulkTransferStatusRetrying             BulkTransferStatus = 13
	BulkTransferStatusCompleted            BulkTransferStatus = 20
	BulkTransferStatusPartiallyUnavailable BulkTransferStatus = 30
	BulkTransferStatusFailed               BulkTransferStatus = 40
)

type BulkTransferItemStatus int

const (
	BulkTransferItemStatusSuccess BulkTransferItemStatus = 1
	BulkTransferItemStatusFailure BulkTransferItemStatus = 2
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

type ClientSecretType int

const (
	ClientSecretTypeBasic ClientSecretType = 1
	ClientSecretTypePost  ClientSecretType = 2
)
