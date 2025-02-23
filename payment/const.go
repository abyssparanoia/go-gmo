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
	secureTran2Path               = "payment/SecureTran2.idPass"
	tds2ResultPath                = "payment/Tds2Result.idPass"
	saveMemberPath                = "payment/SaveMember.idPass"
	updateMemberPath              = "payment/UpdateMember.idPass"
	deleteMemberPath              = "payment/DeleteMember.idPass"
	searchMemberPath              = "payment/SearchMember.idPass"
	saveCardPath                  = "payment/SaveCard.idPass"
	tradedCardPath                = "payment/TradedCard.idPass"
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
	payPayExecTranPath            = "payment/ExecTranPaypay.idPass"
	payPaySalesPath               = "payment/PaypaySales.idPass"
	payPaySCancelReturnPath       = "payment/PaypayCancelReturn.idPass"
	postpayEntryTranPath          = "payment/EntryTranPostpay.idPass"
	postpayExecTranPath           = "payment/ExecTranPostpay.idPass"
	postpayShippedTranPath        = "payment/PostpayShipping.idPass"
	postpayChangeTranPath         = "payment/PostpayChange.idPass"
	postpayCancelTranPath         = "payment/PostpayCancel.idPass"
	linkPlusGetUrlPaymentPath     = "payment/GetLinkplusUrlPayment.json"
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
	// TradeMultiStatusReqSuccess ... request success
	TradeMultiStatusReqSuccess TradeMultiStatus = "REQSUCCESS"
	// TradeMultiStatusAuthProcess ... auth process
	TradeMultiStatusAuthProcess TradeMultiStatus = "AUTHPROCESS"
	// TradeMultiStatusAuthenticated ... authenticated
	TradeMultiStatusAuthenticated TradeMultiStatus = "AUTHENTICATED"
	// TradeMultiStatusAuth ... auth
	TradeMultiStatusAuth TradeMultiStatus = "AUTH"
	// TradeMultiStatusTrading ... proccessing
	TradeMultiStatusTrading TradeMultiStatus = "TRADING"
	// TradeMultiStatusPaysuccess ... success
	TradeMultiStatusPaysuccess TradeMultiStatus = "PAYSUCCESS"
	// TradeMultiStatusReqSales ... request sales
	TradeMultiStatusReqSales TradeMultiStatus = "REQSALES"
	// TradeMultiStatusSales ... sales
	TradeMultiStatusSales TradeMultiStatus = "SALES"
	// TradeMultiStatusCapture ... capture
	TradeMultiStatusCapture TradeMultiStatus = "CAPTURE"
	// TradeMultiStatusCancel ... cancel
	TradeMultiStatusCancel TradeMultiStatus = "CANCEL"
	// TradeMultiStatusReturn ... return
	TradeMultiStatusReturn TradeMultiStatus = "RETURN"
	// TradeMultiStatusPayFail ... pay fail
	TradeMultiStatusPayFail TradeMultiStatus = "PAYFAIL"
	// TradeMultiStatusStop ... stop
	TradeMultiStatusStop TradeMultiStatus = "STOP"
	// TradeMultiStatusExpired ... expired
	TradeMultiStatusExpired TradeMultiStatus = "EXPIRED"
	// TradeMultiStatusShipped ... shipped
	TradeMultiStatusShipped TradeMultiStatus = "SHIPPED"
	// TradeMultiStatusInvoice ... invoice
	TradeMultiStatusInvoice TradeMultiStatus = "INVOICE"
	// TradeMultiStatusForceCancel ... force cancel
	TradeMultiStatusForceCancel TradeMultiStatus = "FORCECANCEL"
)

func (s TradeMultiStatus) String() string {
	return string(s)
}

// PayType ... pay type
type PayType string

const (
	CreditCardPayType PayType = "0"
	// PayEasyPayType ... convenience store
	CvsPayType PayType = "3"
	// PayEasyPayType ... pay easy
	PayEasyPayType PayType = "4"
	// IDNetPayType ... iD net
	IDNetPayType PayType = "6"
	// GANBPayType ... GMO Aozora Net Bank
	GANBPayType PayType = "36"
	// PayPayPayType ... PayPay
	PostpayPayType PayType = "44"
	// PayPayPayType ... PayPay
	PayPayPayType PayType = "45"
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

type PayPayTransitionType int

const (
	PayPayTransitionTypeWebLink  PayPayTransitionType = 1
	PayPayTransitionTypeDeepLink PayPayTransitionType = 2
)

type PostpayCustomerPaymentType int

const (
	PostpayCustomerPaymentTypeInvoiceSentSeparately PostpayCustomerPaymentType = 2
	PostpayCustomerPaymentTypeInvoiceIncluded       PostpayCustomerPaymentType = 3
)

type PostpayCustomerSex int

const (
	PostpayCustomerSexUnknown PostpayCustomerSex = 0
	PostpayCustomerSexMale    PostpayCustomerSex = 1
	PostpayCustomerSexFemale  PostpayCustomerSex = 2
)

type PostpayPDCompanyCode int

const (
	PostpayPDCompanyCodeUnknown    PostpayPDCompanyCode = 0
	PostpayPDCompanyCodeSagawa     PostpayPDCompanyCode = 11
	PostpayPDCompanyCodeYamato     PostpayPDCompanyCode = 12
	PostpayPDCompanyCodeNishina    PostpayPDCompanyCode = 14
	PostpayPDCompanyCodeKakitome   PostpayPDCompanyCode = 15
	PostpayPDCompanyCodeYuPack     PostpayPDCompanyCode = 16
	PostpayPDCompanyCodeFukuyama   PostpayPDCompanyCode = 17
	PostpayPDCompanyCodeEcoHai     PostpayPDCompanyCode = 27
	PostpayPDCompanyCodeLetterPack PostpayPDCompanyCode = 28
)

type LinkPlusPayMethod string

const (
	LinkPlusPayMethodCredit         LinkPlusPayMethod = "credit"
	LinkPlusPayMethodCvs            LinkPlusPayMethod = "cvs"
	LinkPlusPayMethodPayEasy        LinkPlusPayMethod = "payeasy"
	LinkPlusPayMethodDocomo         LinkPlusPayMethod = "docomo"
	LinkPlusPayMethodAu             LinkPlusPayMethod = "au"
	LinkPlusPayMethodSb             LinkPlusPayMethod = "sb"
	LinkPlusPayMethodEpos           LinkPlusPayMethod = "epospay"
	LinkPlusPayMethodDcc            LinkPlusPayMethod = "dcc"
	LinkPlusPayMethodLinePay        LinkPlusPayMethod = "linepay"
	LinkPlusPayMethodFamiPay        LinkPlusPayMethod = "famipay"
	LinkPlusPayMethodMerPay         LinkPlusPayMethod = "merpay"
	LinkPlusPayMethodRakutenID      LinkPlusPayMethod = "rakutenid"
	LinkPlusPayMethodRakutenPayV2   LinkPlusPayMethod = "rakutenpayv2"
	LinkPlusPayMethodPaypay         LinkPlusPayMethod = "paypay"
	LinkPlusPayMethodVirtualAccount LinkPlusPayMethod = "virtualaccount"
	LinkPlusPayMethodAuPay          LinkPlusPayMethod = "aupay"
	LinkPlusPayMethodGanb           LinkPlusPayMethod = "ganb"
	LinkPlusPayMethodUnionPay       LinkPlusPayMethod = "unionpay"
)

type LinkPlusTemplateID string

const (
	LinkPlusTemplateIDDesignA LinkPlusTemplateID = "designA"
	LinkPlusTemplateIDDesignB LinkPlusTemplateID = "designB"
	LinkPlusTemplateIDDesignC LinkPlusTemplateID = "designC"
	LinkPlusTemplateIDDesignD LinkPlusTemplateID = "designD"
)

type LinkPlusColorPattern string

const (
	LinkPlusColorPatternBlue      LinkPlusColorPattern = "blue_01"
	LinkPlusColorPatternBlueGray  LinkPlusColorPattern = "bluegray_01"
	LinkPlusColorPatternSkyBlue   LinkPlusColorPattern = "skyblue_01"
	LinkPlusColorPatternGreen     LinkPlusColorPattern = "pink_01"
	LinkPlusColorPatternYellow    LinkPlusColorPattern = "yellow_01"
	LinkPlusColorPatternBlack     LinkPlusColorPattern = "black_01"
	LinkPlusColorPatternNature    LinkPlusColorPattern = "nature_01"
	LinkPlusColorPatternGreenGray LinkPlusColorPattern = "greengray_01"
)

type LinkPlusLang string

const (
	LinkPlusLangJP LinkPlusLang = "ja"
	LinkPlusLangEN LinkPlusLang = "en"
	LinkPlusLangZH LinkPlusLang = "zh"
)

type SecureTran2CallbackType string

const (
	SecureTran2CallbackTypePost SecureTran2CallbackType = "1"
	SecureTran2CallbackTypeGet  SecureTran2CallbackType = "3"
)

type TDFlag int

const (
	TDFlagEnabled TDFlag = 2
)

type TDS2ChallengeIndType string

const (
	TDS2ChallengeIndTypeRequired    TDS2ChallengeIndType = "1"
	TDS2ChallengeIndTypeNotRequired TDS2ChallengeIndType = "2"
)

type TDS2TransResult string

const (
	TDS2TransResultY TDS2TransResult = "Y" // 認証／カード番号確認に成功
	TDS2TransResultA TDS2TransResult = "A" // 処理の試行が実施された
	TDS2TransResultN TDS2TransResult = "N" // 未認証／カード番号未確認。取引拒否
	TDS2TransResultU TDS2TransResult = "U" // 認証／カード番号確認を実行できなかった
	TDS2TransResultR TDS2TransResult = "R" // 認証 / カード番号確認が拒否された
)

type TDS2TransResultReason string

const (
	TDS2TransResultReasonCardAuthenticationFailed             TDS2TransResultReason = "01" // カード認証に失敗した
	TDS2TransResultReasonUnknownDevice                        TDS2TransResultReason = "02" // 不明なデバイス
	TDS2TransResultReasonUnsupportedDevice                    TDS2TransResultReason = "03" // サポートされていないデバイス
	TDS2TransResultReasonAuthenticationFrequencyExceeded      TDS2TransResultReason = "04" // 認証頻度の上限を超えた
	TDS2TransResultReasonExpiredCard                          TDS2TransResultReason = "05" // 有効期限切れのカード
	TDS2TransResultReasonInvalidCardNumber                    TDS2TransResultReason = "06" // 無効なカード番号
	TDS2TransResultReasonInvalidTransaction                   TDS2TransResultReason = "07" // 無効な取引
	TDS2TransResultReasonInvalidRecord                        TDS2TransResultReason = "08" // 無効なレコード
	TDS2TransResultReasonTechnicalProblem                     TDS2TransResultReason = "09" // 技術的な問題
	TDS2TransResultReasonStolenCard                           TDS2TransResultReason = "10" // 盗難カード
	TDS2TransResultReasonFraudSuspicion                       TDS2TransResultReason = "11" // 不正の疑い
	TDS2TransResultReasonCardMemberNotPermitted               TDS2TransResultReason = "12" // カード会員に取引が許可されていない
	TDS2TransResultReasonCardMemberNotRegistered              TDS2TransResultReason = "13" // カード会員がサービスに登録されていない
	TDS2TransResultReasonACSExpired                           TDS2TransResultReason = "14" // ACSの最大チャレンジを超える
	TDS2TransResultReasonLowTrust                             TDS2TransResultReason = "15" // 信頼度が低い
	TDS2TransResultReasonMediumTrust                          TDS2TransResultReason = "16" // 信頼度が中程度
	TDS2TransResultReasonHighTrust                            TDS2TransResultReason = "17" // 信頼度が高い
	TDS2TransResultReasonVeryHighTrust                        TDS2TransResultReason = "18" // 信頼度が非常に高い
	TDS2TransResultReasonACSMaxChallengeExceeded              TDS2TransResultReason = "19" // ACSの最大チャレンジを超える
	TDS2TransResultReasonNonSettlementTransactionNotSupported TDS2TransResultReason = "20" // 非決済取引はサポートされていません
	TDS2TransResultReason3RITransactionNotSupported           TDS2TransResultReason = "21" // 3RIトランザクションはサポートされていません
	TDS2TransResultReasonACSTechnicalProblem                  TDS2TransResultReason = "22" // ACSの技術的な問題
	TDS2TransResultReasonAuthenticationNotPerformed           TDS2TransResultReason = "26" // 認証は試行されましたが、カード会員によって実行されませんでした
	TDS2TransResultReasonOther                                TDS2TransResultReason = "80" // その他
)
