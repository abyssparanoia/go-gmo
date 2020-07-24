package deferred

import (
	"context"
	"errors"
)

var (
	errInvalidParameterPassed = errors.New("invalid parameter passed")
)

type shopInfo struct {
	AuthenticationID string `xml:"authenticationId"`
	ShopCode         string `xml:"shopCode"`
	ConnectPassword  string `xml:"connectPassword"`
}

type httpInfo struct {
	HTTPHeader string `xml:"httpHeader"`
	DeviceInfo string `xml:"deviceInfo"`
}

type registerRequestParam struct {
	ShopInfo   *shopInfo  `xml:"shopInfo"`
	HTTPInfo   *httpInfo  `xml:"httpInfo"`
	Buyer      *buyer     `xml:"buyer"`
	Deliveries deliveries `xml:"deliveries"`
}

type buyer struct {
	GMOTransactionID  string `xml:"gmoTransactionId"`
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
	DetailName     string `xml:"detailName"`
	DetailPrice    string `xml:"detailPrice"`
	DetailQuantity string `xml:"detailQuantity"`
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
	body := req.toParam()
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
