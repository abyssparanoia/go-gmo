package deferred

import (
	"context"
	"errors"
)

var (
	errInvalidParameterPassed = errors.New("invalid parameter passed")
)

type shopInfo struct {
	AuthenticationID string `xml:"authenticationId" validate:"required"`
	ShopCode         string `xml:"shopCode" validate:"required"`
	ConnectPassword  string `xml:"connectPassword" validate:"required"`
}

type httpInfo struct {
	HTTPHeader string `xml:"httpHeader"`
	DeviceInfo string `xml:"deviceInfo"`
}

type registerRequestParam struct {
	ShopInfo   *shopInfo  `xml:"shopInfo"`
	HTTPInfo   *httpInfo  `xml:"httpInfo"`
	Buyer      *buyer     `xml:"buyer" validate:"required"`
	Deliveries deliveries `xml:"deliveries" validate:"required"`
}

type buyer struct {
	GMOTransactionID  string `xml:"gmoTransactionId"`
	ShopTransactionID string `xml:"shopTransactionId" validate:"required"`
	ShopOrderDate     string `xml:"shopOrderDate" validate:"required"`
	FullName          string `xml:"fullName" validate:"required"`
	FullNameKana      string `xml:"fullKanaName"`
	ZipCode           string `xml:"zipCode" validate:"required"`
	Address           string `xml:"address" validate:"required"`
	CompanyName       string `xml:"companyName"`
	DepartmentName    string `xml:"departmentName"`
	Tel1              string `xml:"tel1" validate:"required"`
	Tel2              string `xml:"tel2"`
	Email             string `xml:"email1" validate:"required"`
	Email2            string `xml:"email2"`
	BilledAmount      string `xml:"billedAmount" validate:"required"`
	GMOExtend1        string `xml:"gmoExtend1"`
	PaymentType       string `xml:"paymentType" validate:"required"`
	Sex               string `xml:"sex"`
	BirthDay          string `xml:"birthDay"`
	MemberRegistDate  string `xml:"memberRegistDate"`
	BuyCount          string `xml:"buyCount"`
	BuyAmountTotal    string `xml:"buyAmoutTotal"`
	MemberID          string `xml:"memberId"`
}

type buyerModification struct {
	GMOTransactionID  string `xml:"gmoTransactionId" validate:"required"`
	ShopTransactionID string `xml:"shopTransactionId"`
	ShopOrderDate     string `xml:"shopOrderDate"`
	FullName          string `xml:"fullName"`
	FullNameKana      string `xml:"fullKanaName"`
	ZipCode           string `xml:"zipCode"`
	Address           string `xml:"address"`
	CompanyName       string `xml:"companyName"`
	DepartmentName    string `xml:"departmentName"`
	Tel1              string `xml:"tel1"`
	Tel2              string `xml:"tel2"`
	Email             string `xml:"email1"`
	Email2            string `xml:"email2"`
	BilledAmount      string `xml:"billedAmount"`
	GMOExtend1        string `xml:"gmoExtend1"`
	PaymentType       string `xml:"paymentType"`
	Sex               string `xml:"sex"`
	BirthDay          string `xml:"birthDay"`
	MemberRegistDate  string `xml:"memberRegistDate"`
	BuyCount          string `xml:"buyCount"`
	BuyAmountTotal    string `xml:"buyAmoutTotal"`
	MemberID          string `xml:"memberId"`
}

func (o *Buyer) toParamModification() *buyerModification {
	if o == nil {
		return nil
	}
	p := &buyerModification{
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

type delivery struct {
	DeliveryCustomer *deliveryCustomer `xml:"deliveryCustomer"`
	Details          details           `xml:"details"`
}

type deliveries struct {
	DeliveryInner []*delivery `xml:"delivery"`
}

type deliveryCustomer struct {
	FullName       string `xml:"fullName"`
	FullNameKana   string `xml:"fullKanaName"`
	ZipCode        string `xml:"zipCode"`
	Address        string `xml:"address"`
	CompanyName    string `xml:"companyName"`
	DepartmentName string `xml:"departmentName"`
	Tel            string `xml:"tel"`
}

type detail struct {
	DetailName     string `xml:"detailName" validate:"required"`
	DetailPrice    string `xml:"detailPrice" validate:"required"`
	DetailQuantity string `xml:"detailQuantity" validate:"required"`
	GMOExtend2     string `xml:"gmoExtend2"`
	GMOExtend3     string `xml:"gmoExtend3"`
	GMOExtend4     string `xml:"gmoExtend4"`
	DetailBrand    string `xml:"detailBrand"`
	DetailCategory string `xml:"detailCategory"`
}

type details struct {
	DetailsInner []*detail `xml:"detail"`
}

type registerResponseParam struct {
	Result            string             `xml:"result"`
	Errors            *gmoErrors         `xml:"errors"`
	TransactionResult *transactionResult `xml:"transactionResult"`
}

type gmoError struct {
	ErrorCode    string `xml:"errorCode"`
	ErrorMessage string `xml:"errorMessage"`
}

type gmoErrors struct {
	ErrorsInner []*gmoError `xml:"error"`
}

type transactionResult struct {
	ShopTransactionID string   `xml:"shopTransactionId"`
	GMOTransactionID  string   `xml:"gmoTransactionId"`
	AuthorResult      string   `xml:"authorResult"`
	AutoAutherResult  string   `xml:"autoAuthorResult"`
	MaulAuthorResult  string   `xml:"maulAuthorResult"`
	Reasons           []string `xml:"reasons"`
}

func (c *Client) RegisterTransaction(ctx context.Context, req *RegisterRequestParam) (*RegisterResponseParam, error) {
	if req == nil {
		return nil, errInvalidParameterPassed
	}
	body, err := req.toParam()
	if err != nil {
		return nil, err
	}
	respParam := registerResponseParam{}
	body.ShopInfo = &shopInfo{
		AuthenticationID: c.AuthenticationID,
		ShopCode:         c.ShopCode,
		ConnectPassword:  c.ConnectPassword,
	}
	status, err := c.doAndUnmarshalXML(ctx, POST, c.APIHost, []string{"auto", "transaction.do"}, map[string]string{},
		body, &respParam)
	if err != nil {
		return nil, err
	}
	resp := newRegisterResponseParam(&respParam)
	resp.Status = status
	return resp, nil
}
