package deferred

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

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
	Errors                        Errors
}

func newInvoiceGetResponse(o *invoiceResult) *InvoiceGetResponse {
	if o == nil {
		return nil
	}
	p := &InvoiceGetResponse{
		GMOTransactionID:              o.InvoiceDataResult.GMOTransactionID,
		DeliveryZip:                   o.InvoiceDataResult.DeliveryZip,
		DeliveryAddress1:              o.InvoiceDataResult.DeliveryAddress1,
		DeliveryAddress2:              o.InvoiceDataResult.DeliveryAddress2,
		PurchaseCompanyName:           o.InvoiceDataResult.PurchaseCompanyName,
		PurchaseDepartmentName:        o.InvoiceDataResult.PurchaseDepartmentName,
		PurchaseUserName:              o.InvoiceDataResult.PurchaseUserName,
		ShopName:                      o.InvoiceDataResult.ShopName,
		ShopTransactionID:             o.InvoiceDataResult.ShopTransactionID,
		InvoiceMatter1:                o.InvoiceDataResult.InvoiceMatter1,
		InvoiceMatter2:                o.InvoiceDataResult.InvoiceMatter2,
		InvoiceMatter3:                o.InvoiceDataResult.InvoiceMatter3,
		InvoiceMatter4:                o.InvoiceDataResult.InvoiceMatter4,
		InvoiceMatter5:                o.InvoiceDataResult.InvoiceMatter5,
		GMOCompanyName:                o.InvoiceDataResult.GMOCompanyName,
		GMOInfo1:                      o.InvoiceDataResult.GMOInfo1,
		GMOInfo2:                      o.InvoiceDataResult.GMOInfo2,
		GMOInfo3:                      o.InvoiceDataResult.GMOInfo3,
		GMOInfo4:                      o.InvoiceDataResult.GMOInfo4,
		InvoiceTitle:                  o.InvoiceDataResult.InvoiceTitle,
		InvoiceGreeting1:              o.InvoiceDataResult.InvoiceGreeting1,
		InvoiceGreeting2:              o.InvoiceDataResult.InvoiceGreeting2,
		InvoiceGreeting3:              o.InvoiceDataResult.InvoiceGreeting3,
		InvoiceGreeting4:              o.InvoiceDataResult.InvoiceGreeting4,
		Yobi1:                         o.InvoiceDataResult.Yobi1,
		Yobi2:                         o.InvoiceDataResult.Yobi2,
		Yobi3:                         o.InvoiceDataResult.Yobi3,
		Yobi4:                         o.InvoiceDataResult.Yobi4,
		Yobi5:                         o.InvoiceDataResult.Yobi5,
		Yobi6:                         o.InvoiceDataResult.Yobi6,
		Yobi7:                         o.InvoiceDataResult.Yobi7,
		Yobi8:                         o.InvoiceDataResult.Yobi8,
		Yobi9:                         o.InvoiceDataResult.Yobi9,
		Yobi10:                        o.InvoiceDataResult.Yobi10,
		BilledAmount:                  o.InvoiceDataResult.BilledAmount,
		BilledAmountTax:               o.InvoiceDataResult.BilledAmountTax,
		OrderDate:                     o.InvoiceDataResult.OrderDate,
		InvoiceIssueDate:              o.InvoiceDataResult.InvoiceIssueDate,
		PaymentDueDate:                o.InvoiceDataResult.PaymentDueDate,
		TrackingNumber:                o.InvoiceDataResult.TrackingNumber,
		BankNoteWording:               o.InvoiceDataResult.BankNoteWording,
		BankName:                      o.InvoiceDataResult.BankName,
		Bankcode:                      o.InvoiceDataResult.Bankcode,
		DepositType:                   o.InvoiceDataResult.DepositType,
		AccountNumber:                 o.InvoiceDataResult.AccountNumber,
		BankAccountHolder:             o.InvoiceDataResult.BankAccountHolder,
		VotesBilledAmount:             o.InvoiceDataResult.VotesBilledAmount,
		VotesFrontUpperInfo:           o.InvoiceDataResult.VotesFrontUpperInfo,
		VotesBarCode:                  o.InvoiceDataResult.VotesBarCode,
		DocketBilledAmount:            o.InvoiceDataResult.DocketBilledAmount,
		DocketPurchaseAddress:         o.InvoiceDataResult.DocketPurchaseAddress,
		DocketPurchaseCompanyName:     o.InvoiceDataResult.DocketPurchaseCompanyName,
		DocketPurchaseDepartmentName:  o.InvoiceDataResult.DocketPurchaseDepartmentName,
		DocketPurchaseUserName:        o.InvoiceDataResult.DocketPurchaseUserName,
		DocketTrackingNumber:          o.InvoiceDataResult.DocketTrackingNumber,
		DocketX:                       o.InvoiceDataResult.DocketX,
		ReceiptPurchaseCompanyName:    o.InvoiceDataResult.ReceiptPurchaseCompanyName,
		ReceiptPurchaseDepartmentName: o.InvoiceDataResult.ReceiptPurchaseDepartmentName,
		ReceiptPurchaseUserName:       o.InvoiceDataResult.ReceiptPurchaseUserName,
		ReceiptTrackingNumber1:        o.InvoiceDataResult.ReceiptTrackingNumber1,
		ReceiptTrackingNumber2:        o.InvoiceDataResult.ReceiptTrackingNumber2,
		ReceiptAmount:                 o.InvoiceDataResult.ReceiptAmount,
		ReceiptTax:                    o.InvoiceDataResult.ReceiptTax,
		ReceiptShopName:               o.InvoiceDataResult.ReceiptShopName,
		ReceiptPrintWord:              o.InvoiceDataResult.ReceiptPrintWord,
		String:                        o.InvoiceDataResult.String,
		Yobi11:                        o.InvoiceDataResult.Yobi11,
		Yobi12:                        o.InvoiceDataResult.Yobi12,
		Yobi13:                        o.InvoiceDataResult.Yobi13,
		Yobi14:                        o.InvoiceDataResult.Yobi14,
		Yobi15:                        o.InvoiceDataResult.Yobi15,
		GoodsDetail: func() GoodsDetail {
			return GoodsDetail{
				GoodsName:   o.InvoiceDataResult.DetailList.GoodsDetail.GoodsName,
				GoodsNum:    o.InvoiceDataResult.DetailList.GoodsDetail.GoodsNum,
				GoodsPrice:  o.InvoiceDataResult.DetailList.GoodsDetail.GoodsPrice,
				GoodsAmount: o.InvoiceDataResult.DetailList.GoodsDetail.GoodsAmount,
				GoodsTax:    o.InvoiceDataResult.DetailList.GoodsDetail.GoodsTax,
				Yobi16:      o.InvoiceDataResult.DetailList.GoodsDetail.Yobi16,
				Yobi17:      o.InvoiceDataResult.DetailList.GoodsDetail.Yobi17,
				Yobi18:      o.InvoiceDataResult.DetailList.GoodsDetail.Yobi18,
				Yobi19:      o.InvoiceDataResult.DetailList.GoodsDetail.Yobi19,
				Yobi20:      o.InvoiceDataResult.DetailList.GoodsDetail.Yobi20,
			}
		}(),
		Errors: func() Errors {
			if o.Errors == nil {
				return nil
			}
			r := make(Errors, len(o.Errors.ErrorsInner))
			for i, d := range o.Errors.ErrorsInner {
				r[i] = newInvoiceError(d)
			}
			return r
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
