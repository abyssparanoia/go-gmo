package deferred

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

func (o *RegisterRequestParam) toParam() *registerRequestParam {
	if o == nil {
		return nil
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
	return p
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
	p := &RegisterResponseParam{
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

type Error struct {
	ErrorCode    string
	ErrorMessage string
}

func newError(o *gmoError) *Error {
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

func (o *ModifyRequest) toParam() *modifyRequest {
	if o == nil {
		return nil
	}
	p := &modifyRequest{
		Buyer: o.Buyer.toParam(),
		Deliveries: func() deliveries {
			r := make([]*delivery, len(o.Deliveries))
			for i, d := range o.Deliveries {
				r[i] = d.toParam()
			}
			return deliveries{r}
		}(),
		KindInfo: o.KindInfo.toParam(),
	}
	return p
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

func (o *AuthResultGetRequest) toParam() *authResultGetRequest {
	if o == nil {
		return nil
	}
	p := &authResultGetRequest{
		Transaction: o.Transaction.toParam(),
	}
	return p
}

type AuthResultGetResponse struct {
	Result            string
	Errors            Errors
	TransactionResult *TransactionResult
	Status            int
}

func newAuthResultGetResponse(o *authResultGetResponse) *AuthResultGetResponse {
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

func (o *ShippingReportRequest) toParam() *shippingReportRequest {
	if o == nil {
		return nil
	}
	p := &shippingReportRequest{
		Transaction: o.Transaction.toParam(),
	}
	return p
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

func (o *ShippingModifyRequest) toParam() *shippingModifyRequest {
	if o == nil {
		return nil
	}
	p := &shippingModifyRequest{
		Transaction: o.Transaction.toParam(),
		KindInfo:    o.KindInfo.toParam(),
	}
	return p
}

type ShippingModifyResponse struct {
	Result            string
	Errors            Errors
	Status            int
	TransactionResult *TransactionResult
}

func newShippingModifyResponse(o *shippingModifyResponse) *ShippingModifyResponse {
	p := &ShippingModifyResponse{
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
		TransactionResult: func() *TransactionResult {
			if o.TransactionResult == nil {
				return &TransactionResult{}
			}
			return newTransactionResult(o.TransactionResult)
		}(),
	}
	return p
}
