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
	GMOTransactionID  string `xml:"gmoTransactionId,omitempty"`
	ShopTransactionID string `xml:"shopTransactionId" validate:"required"`
	ShopOrderDate     string `xml:"shopOrderDate" validate:"required"`
	FullName          string `xml:"fullName" validate:"required"`
	FullNameKana      string `xml:"fullKanaName,omitempty"`
	ZipCode           string `xml:"zipCode" validate:"required"`
	Address           string `xml:"address" validate:"required"`
	CompanyName       string `xml:"companyName,omitempty"`
	DepartmentName    string `xml:"departmentName,omitempty"`
	Tel1              string `xml:"tel1" validate:"required"`
	Tel2              string `xml:"tel2,omitempty"`
	Email             string `xml:"email1" validate:"required"`
	Email2            string `xml:"email2,omitempty"`
	BilledAmount      string `xml:"billedAmount,omitempty" validate:"required"`
	GMOExtend1        string `xml:"gmoExtend1,omitempty"`
	PaymentType       string `xml:"paymentType,omitempty" validate:"required"`
	Sex               string `xml:"sex,omitempty"`
	BirthDay          string `xml:"birthDay,omitempty"`
	MemberRegistDate  string `xml:"memberRegistDate,omitempty"`
	BuyCount          string `xml:"buyCount,omitempty"`
	BuyAmountTotal    string `xml:"buyAmoutTotal,omitempty"`
	MemberID          string `xml:"memberId,omitempty"`
}

type delivery struct {
	DeliveryCustomer *deliveryCustomer `xml:"deliveryCustomer"`
	Details          details           `xml:"details"`
}

type deliveries struct {
	DeliveryInner []*delivery `xml:"delivery"`
}

type deliveryCustomer struct {
	FullName       string `xml:"fullName,omitempty"`
	FullNameKana   string `xml:"fullKanaName,omitempty"`
	ZipCode        string `xml:"zipCode,omitempty"`
	Address        string `xml:"address,omitempty"`
	CompanyName    string `xml:"companyName,omitempty"`
	DepartmentName string `xml:"departmentName,omitempty"`
	Tel            string `xml:"tel,omitempty"`
}

type detail struct {
	DetailName     string `xml:"detailName" validate:"required"`
	DetailPrice    string `xml:"detailPrice" validate:"required"`
	DetailQuantity string `xml:"detailQuantity" validate:"required"`
	GMOExtend2     string `xml:"gmoExtend2,omitempty"`
	GMOExtend3     string `xml:"gmoExtend3,omitempty"`
	GMOExtend4     string `xml:"gmoExtend4,omitempty"`
	DetailBrand    string `xml:"detailBrand,omitempty"`
	DetailCategory string `xml:"detailCategory,omitempty"`
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
