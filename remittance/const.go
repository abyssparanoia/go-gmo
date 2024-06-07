package remittance

const (
	apiHostSandbox              = "https://test-remittance.gmopg.jp"
	apiHostProduction           = "https://remittance.gmopg.jp"
	apiHostTest                 = "http://remittance.gmopg.jp"
	mailDepositRegistrationPath = "api/shop/MailDepositRegistration.json"
	accountRegistrationPath     = "api/AccountRegistration.json"
	depositRegistrationPath     = "api/DepositRegistration.json"
	realDepositRegistrationPath = "api/RealDepositRegistration.json"
)

type MailDepositMethod string

const (
	MailDepositMethodRegister MailDepositMethod = "1"
	MailDepositMethodCancel   MailDepositMethod = "2"
)

type SelectablePaymentMethod string

const (
	SelectablePaymentMethodDisable SelectablePaymentMethod = "0"
	SelectablePaymentMethodEnable  SelectablePaymentMethod = "1"
)

type BankAccountRegistrationMethod string

const (
	BankAccountRegistrationMethodRegister BankAccountRegistrationMethod = "1"
	BankAccountRegistrationMethodUpdate   BankAccountRegistrationMethod = "2"
	BankAccountRegistrationMethodDelete   BankAccountRegistrationMethod = "3"
)

type DepositRegistrationMethod string

const (
	DepositRegistrationMethodRegister DepositRegistrationMethod = "1"
	DepositRegistrationMethodDelete   DepositRegistrationMethod = "2"
)

type DepositResult string

const (
	DepositResultCompleted DepositResult = "1"
	DepositResultFailed    DepositResult = "4"
)
