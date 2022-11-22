package remittance

const (
	apiHostSandbox              = "https://test-remittance.gmopg.jp"
	apiHostProduction           = "https://remittance.gmopg.jp"
	apiHostTest                 = "http://remittance.gmopg.jp"
	mailDepositRegistrationPath = "api/shop/MailDepositRegistration.json"
)

type Method string

const (
	MethodRegister Method = "1"
	MethodCancel   Method = "2"
)

type SelectablePaymentMethod string

const (
	SelectablePaymentMethodDisable SelectablePaymentMethod = "0"
	SelectablePaymentMethodEnable  SelectablePaymentMethod = "1"
)
