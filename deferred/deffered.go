package deferred

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// structs for library client

type UpdateKind uint8

func (u UpdateKind) Uint8() uint8 {
	return uint8(u)
}

const (
	_ UpdateKind = iota
	Modification
	Cancel
)

type RegisterRequestParam struct {
	Buyer      *Buyer
	Deliveries Deliveries
}

func (o *RegisterRequestParam) toParam() (*registerRequestParam, error) {
	if o == nil {
		return nil, nil
	}
	p := &registerRequestParam{
		Buyer: o.Buyer.toParam(),
		Deliveries: func() deliveries {
			r := make([]*delivery, len(o.Deliveries))
			for i, d := range o.Deliveries {
				r[i] = d.toParam()
			}
			return deliveries{r}
		}(),
	}
	if err := validate.Struct(p); err != nil {
		return nil, err
	}
	return p, nil
}

type Buyer struct {
	GMOTransactionID  string
	ShopTransactionID string
	ShopOrderDate     string
	FullName          string
	FullNameKana      string
	ZipCode           string
	Address           string
	CompanyName       string
	DepartmentName    string
	Tel1              string
	Tel2              string
	Email             string
	Email2            string
	BilledAmount      string
	GMOExtend1        string
	PaymentType       string
	Sex               string
	BirthDay          string
	MemberRegistDate  string
	BuyCount          string
	BuyAmountTotal    string
	MemberID          string
}

func (o *Buyer) toParam() *buyer {
	if o == nil {
		return nil
	}
	p := &buyer{
		GMOTransactionID:  o.GMOTransactionID,
		ShopTransactionID: o.ShopTransactionID,
		ShopOrderDate:     o.ShopOrderDate,
		FullName:          o.FullName,
		FullNameKana:      o.FullNameKana,
		ZipCode:           o.ZipCode,
		Address:           o.Address,
		CompanyName:       o.CompanyName,
		DepartmentName:    o.DepartmentName,
		Tel1:              o.Tel1,
		Tel2:              o.Tel2,
		Email:             o.Email,
		Email2:            o.Email2,
		BilledAmount:      o.BilledAmount,
		GMOExtend1:        o.GMOExtend1,
		PaymentType:       o.PaymentType,
		Sex:               o.Sex,
		BirthDay:          o.BirthDay,
		MemberRegistDate:  o.MemberRegistDate,
		BuyCount:          o.BuyCount,
		BuyAmountTotal:    o.BuyAmountTotal,
		MemberID:          o.MemberID,
	}
	return p
}

type Delivery struct {
	DeliveryCustomer *DeliveryCustomer
	Details          Details
}

func (o *Delivery) toParam() *delivery {
	if o == nil {
		return nil
	}
	p := &delivery{
		DeliveryCustomer: o.DeliveryCustomer.toParam(),
		Details: func() details {
			r := make([]*detail, len(o.Details))
			for i, d := range o.Details {
				r[i] = d.toParam()
			}
			return details{r}
		}(),
	}
	return p
}

type Deliveries []*Delivery

type DeliveryCustomer struct {
	FullName       string
	FullNameKana   string
	ZipCode        string
	Address        string
	CompanyName    string
	DepartmentName string
	Tel            string
}

func (o *DeliveryCustomer) toParam() *deliveryCustomer {
	if o == nil {
		return nil
	}
	p := &deliveryCustomer{
		FullName:       o.FullName,
		FullNameKana:   o.FullNameKana,
		ZipCode:        o.ZipCode,
		Address:        o.Address,
		CompanyName:    o.CompanyName,
		DepartmentName: o.DepartmentName,
		Tel:            o.Tel,
	}
	return p
}

type Detail struct {
	DetailName     string
	DetailPrice    string
	DetailQuantity string
	GMOExtend2     string
	GMOExtend3     string
	GMOExtend4     string
	DetailBrand    string
	DetailCategory string
}

func (o *Detail) toParam() *detail {
	if o == nil {
		return nil
	}
	p := &detail{
		DetailName:     o.DetailName,
		DetailPrice:    o.DetailPrice,
		DetailQuantity: o.DetailQuantity,
		GMOExtend2:     o.GMOExtend2,
		GMOExtend3:     o.GMOExtend3,
		GMOExtend4:     o.GMOExtend4,
		DetailBrand:    o.DetailBrand,
		DetailCategory: o.DetailCategory,
	}
	return p
}

type Details []*Detail

type RegisterResponseParam struct {
	Result            string
	Errors            Errors
	TransactionResult *TransactionResult
	Status            int
}

func newRegisterResponseParam(o *registerResponseParam) *RegisterResponseParam {
	if o == nil {
		return nil
	}
	p := &RegisterResponseParam{
		Result: o.Result,
		Errors: func() Errors {
			r := make(Errors, len(o.Errors.ErrorsInner))
			for i, d := range o.Errors.ErrorsInner {
				r[i] = newError(d)
			}
			return r
		}(),
		TransactionResult: newTransactionResult(o.TransactionResult),
	}
	return p
}

type Error struct {
	ErrorCode    string
	ErrorMessage string
}

func newError(o *gmoError) *Error {
	if o == nil {
		return nil
	}
	p := &Error{
		ErrorCode:    o.ErrorCode,
		ErrorMessage: o.ErrorMessage,
	}
	return p
}

type Errors []*Error

type TransactionResult struct {
	ShopTransactionID string
	GMOTransactionID  string
	AuthorResult      string
	AutoAutherResult  string
	MaulAuthorResult  string
	Reasons           []string
}

func newTransactionResult(o *transactionResult) *TransactionResult {
	if o == nil {
		return nil
	}
	p := &TransactionResult{
		ShopTransactionID: o.ShopTransactionID,
		GMOTransactionID:  o.GMOTransactionID,
		AuthorResult:      o.AuthorResult,
		AutoAutherResult:  o.AutoAutherResult,
		MaulAuthorResult:  o.MaulAuthorResult,
		Reasons:           o.Reasons,
	}
	return p
}

type ModifyRequest struct {
	Buyer      *Buyer
	Deliveries Deliveries
	KindInfo   *KindInfo
}

func (o *ModifyRequest) toParam() (*modifyRequest, error) {
	if o == nil {
		return nil, nil
	}
	p := &modifyRequest{
		Buyer: o.Buyer.toParamModification(),
		Deliveries: func() deliveries {
			r := make([]*delivery, len(o.Deliveries))
			for i, d := range o.Deliveries {
				r[i] = d.toParam()
			}
			return deliveries{r}
		}(),
		KindInfo: o.KindInfo.toParam(),
	}
	if err := validate.Struct(p); err != nil {
		return nil, err
	}
	return p, nil
}

type KindInfo struct {
	UpdateKind UpdateKind
}

func (o *KindInfo) toParam() *kindInfo {
	if o == nil {
		return nil
	}
	p := &kindInfo{
		UpdateKind: o.UpdateKind.Uint8(),
	}
	return p
}

type ModifyResponse struct {
	Result            string
	Errors            Errors
	TransactionResult *TransactionResult
	Status            int
}

func newModifyResponse(o *modifyResponse) *ModifyResponse {
	if o == nil {
		return nil
	}
	p := &ModifyResponse{
		Result: o.Result,
		Errors: func() Errors {
			if o.Errors == nil {
				return Errors{}
			}
			r := make(Errors, len(o.Errors.ErrorsInner))
			for i, d := range o.Errors.ErrorsInner {
				r[i] = newError(d)
			}
			return r
		}(),
		TransactionResult: newTransactionResult(o.TransactionResult),
	}
	return p
}

type Transaction struct {
	GMOTransactionID string
}

func (o *Transaction) toParam() *transaction {
	if o == nil {
		return nil
	}
	p := &transaction{
		GMOTransactionID: o.GMOTransactionID,
	}
	return p
}

type AuthResultGetRequest struct {
	Transaction *Transaction
}

func (o *AuthResultGetRequest) toParam() (*authResultGetRequest, error) {
	if o == nil {
		return nil, nil
	}
	p := &authResultGetRequest{
		Transaction: o.Transaction.toParam(),
	}
	if err := validate.Struct(p); err != nil {
		return nil, err
	}
	return p, nil
}

type AuthResultGetResponse struct {
	Result            string
	Errors            Errors
	TransactionResult *TransactionResult
	Status            int
}

func newAuthResultGetResponse(o *authResultGetResponse) *AuthResultGetResponse {
	if o == nil {
		return nil
	}
	p := &AuthResultGetResponse{
		Result: o.Result,
		Errors: func() Errors {
			if o.Errors == nil {
				return Errors{}
			}
			r := make(Errors, len(o.Errors.ErrorsInner))
			for i, d := range o.Errors.ErrorsInner {
				r[i] = newError(d)
			}
			return r
		}(),
		TransactionResult: newTransactionResult(o.TransactionResult),
	}
	return p
}

type ShippingReportRequest struct {
	Transaction *ShippingReportTransaction
}

func (o *ShippingReportRequest) toParam() (*shippingReportRequest, error) {
	if o == nil {
		return nil, nil
	}
	p := &shippingReportRequest{
		Transaction: o.Transaction.toParam(),
	}
	if err := validate.Struct(p); err != nil {
		return nil, err
	}
	return p, nil
}

type ShippingReportTransaction struct {
	GMOTransactionID string
	PDCompanyCode    string
	SlipNo           string
}

func (o *ShippingReportTransaction) toParam() *shippingReportTransaction {
	if o == nil {
		return nil
	}
	p := &shippingReportTransaction{
		GMOTransactionID: o.GMOTransactionID,
		PDCompanyCode:    o.PDCompanyCode,
		SlipNo:           o.SlipNo,
	}
	return p
}

type TransactionInfo struct {
	GMOTransactionID string
}

func newTransactionInfo(o *transactionInfo) *TransactionInfo {
	if o == nil {
		return nil
	}
	p := &TransactionInfo{
		GMOTransactionID: o.GMOTransactionID,
	}
	return p
}

type ShippingReportResponse struct {
	Result          string
	Errors          Errors
	Status          int
	TransactionInfo *TransactionInfo
}

func newShippingReportResponse(o *shippingReportResponse) *ShippingReportResponse {
	if o == nil {
		return nil
	}
	p := &ShippingReportResponse{
		Result: o.Result,
		Errors: func() Errors {
			if o.Errors == nil {
				return Errors{}
			}
			r := make(Errors, len(o.Errors.ErrorsInner))
			for i, d := range o.Errors.ErrorsInner {
				r[i] = newError(d)
			}
			return r
		}(),
		TransactionInfo: newTransactionInfo(o.TransactionInfo),
	}
	return p
}

type ShippingModifyRequest struct {
	Transaction *ShippingReportTransaction
	KindInfo    *KindInfo
}

func (o *ShippingModifyRequest) toParam() (*shippingModifyRequest, error) {
	if o == nil {
		return nil, nil
	}
	p := &shippingModifyRequest{
		Transaction: o.Transaction.toParam(),
		KindInfo:    o.KindInfo.toParam(),
	}
	if err := validate.Struct(p); err != nil {
		return nil, err
	}
	return p, nil
}

type ShippingModifyResponse struct {
	Result            string
	Errors            Errors
	Status            int
	TransactionResult *TransactionResult
}

func newShippingModifyResponse(o *shippingModifyResponse) *ShippingModifyResponse {
	if o == nil {
		return nil
	}
	p := &ShippingModifyResponse{
		Result: o.Result,
		Errors: func() Errors {
			r := make(Errors, len(o.Errors.ErrorsInner))
			for i, d := range o.Errors.ErrorsInner {
				r[i] = newError(d)
			}
			return r
		}(),
		TransactionResult: newTransactionResult(o.TransactionResult),
	}
	return p
}

type InvoiceGetRequest struct {
	GMOTransactionID string
}

func (o *InvoiceGetRequest) toParam() (*invoiceRequest, error) {
	p := &invoiceRequest{
		Transaction: &transaction{
			GMOTransactionID: o.GMOTransactionID,
		},
	}
	if err := validate.Struct(p); err != nil {
		return nil, err
	}
	return p, nil
}

type InvoiceGetResponse struct {
	Status                        int
	GMOTransactionID              string
	DeliveryZip                   string
	DeliveryAddress1              string
	DeliveryAddress2              string
	PurchaseCompanyName           string
	PurchaseDepartmentName        string
	PurchaseUserName              string
	ShopName                      string
	ShopTransactionID             string
	InvoiceMatter1                string
	InvoiceMatter2                string
	InvoiceMatter3                string
	InvoiceMatter4                string
	InvoiceMatter5                string
	GMOCompanyName                string
	GMOInfo1                      string
	GMOInfo2                      string
	GMOInfo3                      string
	GMOInfo4                      string
	InvoiceTitle                  string
	InvoiceGreeting1              string
	InvoiceGreeting2              string
	InvoiceGreeting3              string
	InvoiceGreeting4              string
	Yobi1                         string
	Yobi2                         string
	Yobi3                         string
	Yobi4                         string
	Yobi5                         string
	Yobi6                         string
	Yobi7                         string
	Yobi8                         string
	Yobi9                         string
	Yobi10                        string
	BilledAmount                  string
	BilledAmountTax               string
	OrderDate                     string
	InvoiceIssueDate              string
	PaymentDueDate                string
	TrackingNumber                string
	BankNoteWording               string
	BankName                      string
	Bankcode                      string
	DepositType                   string
	AccountNumber                 string
	BankAccountHolder             string
	VotesBilledAmount             string
	VotesFrontUpperInfo           string
	VotesBarCode                  string
	DocketBilledAmount            string
	DocketPurchaseAddress         string
	DocketPurchaseCompanyName     string
	DocketPurchaseDepartmentName  string
	DocketPurchaseUserName        string
	DocketTrackingNumber          string
	DocketX                       string
	ReceiptPurchaseCompanyName    string
	ReceiptPurchaseDepartmentName string
	ReceiptPurchaseUserName       string
	ReceiptTrackingNumber1        string
	ReceiptTrackingNumber2        string
	ReceiptAmount                 string
	ReceiptTax                    string
	ReceiptShopName               string
	ReceiptPrintWord              string
	String                        string
	Yobi11                        string
	Yobi12                        string
	Yobi13                        string
	Yobi14                        string
	Yobi15                        string
	GoodsDetail                   GoodsDetail
}

func newInvoiceGetResponse(o *invoiceResponse) *InvoiceGetResponse {
	p := &InvoiceGetResponse{
		GMOTransactionID:              o.GMOTransactionID,
		DeliveryZip:                   o.DeliveryZip,
		DeliveryAddress1:              o.DeliveryAddress1,
		DeliveryAddress2:              o.DeliveryAddress2,
		PurchaseCompanyName:           o.PurchaseCompanyName,
		PurchaseDepartmentName:        o.PurchaseDepartmentName,
		PurchaseUserName:              o.PurchaseUserName,
		ShopName:                      o.ShopName,
		ShopTransactionID:             o.ShopTransactionID,
		InvoiceMatter1:                o.InvoiceMatter1,
		InvoiceMatter2:                o.InvoiceMatter2,
		InvoiceMatter3:                o.InvoiceMatter3,
		InvoiceMatter4:                o.InvoiceMatter4,
		InvoiceMatter5:                o.InvoiceMatter5,
		GMOCompanyName:                o.GMOCompanyName,
		GMOInfo1:                      o.GMOInfo1,
		GMOInfo2:                      o.GMOInfo2,
		GMOInfo3:                      o.GMOInfo3,
		GMOInfo4:                      o.GMOInfo4,
		InvoiceTitle:                  o.InvoiceTitle,
		InvoiceGreeting1:              o.InvoiceGreeting1,
		InvoiceGreeting2:              o.InvoiceGreeting2,
		InvoiceGreeting3:              o.InvoiceGreeting3,
		InvoiceGreeting4:              o.InvoiceGreeting4,
		Yobi1:                         o.Yobi1,
		Yobi2:                         o.Yobi2,
		Yobi3:                         o.Yobi3,
		Yobi4:                         o.Yobi4,
		Yobi5:                         o.Yobi5,
		Yobi6:                         o.Yobi6,
		Yobi7:                         o.Yobi7,
		Yobi8:                         o.Yobi8,
		Yobi9:                         o.Yobi9,
		Yobi10:                        o.Yobi10,
		BilledAmount:                  o.BilledAmount,
		BilledAmountTax:               o.BilledAmountTax,
		OrderDate:                     o.OrderDate,
		InvoiceIssueDate:              o.InvoiceIssueDate,
		PaymentDueDate:                o.PaymentDueDate,
		TrackingNumber:                o.TrackingNumber,
		BankNoteWording:               o.BankNoteWording,
		BankName:                      o.BankName,
		Bankcode:                      o.Bankcode,
		DepositType:                   o.DepositType,
		AccountNumber:                 o.AccountNumber,
		BankAccountHolder:             o.BankAccountHolder,
		VotesBilledAmount:             o.VotesBilledAmount,
		VotesFrontUpperInfo:           o.VotesFrontUpperInfo,
		VotesBarCode:                  o.VotesBarCode,
		DocketBilledAmount:            o.DocketBilledAmount,
		DocketPurchaseAddress:         o.DocketPurchaseAddress,
		DocketPurchaseCompanyName:     o.DocketPurchaseCompanyName,
		DocketPurchaseDepartmentName:  o.DocketPurchaseDepartmentName,
		DocketPurchaseUserName:        o.DocketPurchaseUserName,
		DocketTrackingNumber:          o.DocketTrackingNumber,
		DocketX:                       o.DocketX,
		ReceiptPurchaseCompanyName:    o.ReceiptPurchaseCompanyName,
		ReceiptPurchaseDepartmentName: o.ReceiptPurchaseDepartmentName,
		ReceiptPurchaseUserName:       o.ReceiptPurchaseUserName,
		ReceiptTrackingNumber1:        o.ReceiptTrackingNumber1,
		ReceiptTrackingNumber2:        o.ReceiptTrackingNumber2,
		ReceiptAmount:                 o.ReceiptAmount,
		ReceiptTax:                    o.ReceiptTax,
		ReceiptShopName:               o.ReceiptShopName,
		ReceiptPrintWord:              o.ReceiptPrintWord,
		String:                        o.String,
		Yobi11:                        o.Yobi11,
		Yobi12:                        o.Yobi12,
		Yobi13:                        o.Yobi13,
		Yobi14:                        o.Yobi14,
		Yobi15:                        o.Yobi15,
		GoodsDetail: func() GoodsDetail {
			return GoodsDetail{
				GoodsName:   o.DetailList.GoodsDetail.GoodsName,
				GoodsNum:    o.DetailList.GoodsDetail.GoodsNum,
				GoodsPrice:  o.DetailList.GoodsDetail.GoodsPrice,
				GoodsAmount: o.DetailList.GoodsDetail.GoodsAmount,
				GoodsTax:    o.DetailList.GoodsDetail.GoodsTax,
				Yobi16:      o.DetailList.GoodsDetail.Yobi16,
				Yobi17:      o.DetailList.GoodsDetail.Yobi17,
				Yobi18:      o.DetailList.GoodsDetail.Yobi18,
				Yobi19:      o.DetailList.GoodsDetail.Yobi19,
				Yobi20:      o.DetailList.GoodsDetail.Yobi20,
			}
		}(),
	}
	return p
}

type GoodsDetail struct {
	GoodsName   string
	GoodsNum    int64
	GoodsPrice  float64
	GoodsAmount float64
	GoodsTax    float64
	Yobi16      string
	Yobi17      string
	Yobi18      string
	Yobi19      string
	Yobi20      string
}
