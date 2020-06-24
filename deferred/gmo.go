package deferred

type RegisterParam struct {
	Buyer Buyer `xml:"buyer"`
}

type Buyer struct {
	ShopTransactionID string `xml:"shopTransactionId"`
	ShopOrderDate     string `xml:"ShopOrderDate"`
	FullName          string `xml:"fullName"`
	FullNameKana      string `xml:"fullNameKana"`
	ZipCode           string `xml:"zipCode"`
	Address           string `xml:"address"`
	CompanyName       string `xml:"companyName"`
	DepartmentName    string `xml:"departmentName"`
	Tel1              string `xml:"tel1"`
	Tel2              string `xml:"tel2"`
	Email             string `xml:"email"`
	Email2            string `xml:"email2"`
	BilledAmount      string `xml:"BilledAmount"`
	GMOExtend1        string `xml:"gmoExtend1"`
	PaymentType       string `xml:"paymentType"`
	Sex               string `xml:"sex"`
	BirthDay          string `xml:"birthDay"`
	MemberRegistDate  string `xml:"memberRegistDate"`
	BuyCount          string `xml:"buyCount"`
	BuyAmountTotal    string `xml:"buyAmountTotal"`
	MemberID          string `xml:"memberID"`
}

type Delivery struct {
	DeliveryCustomer *DeliveryCustomer `xml:"deliveryCustomer"`
	Details          Details           `xml:"details"`
}

type Deliveries []*Delivery

type DeliveryCustomer struct {
	FullName       string `xml:"fullName"`
	FullNameKana   string `xml:"fullNameKana"`
	ZipCode        string `xml:"zipCode"`
	Address        string `xml:"address"`
	CompanyName    string `xml:"companyName"`
	DepartmentName string `xml:"departmentName"`
	Tel            string `xml:"tel"`
}

type Detail struct {
	DetailName     string `xml:"detailName"`
	DetailPrice    string `xml:"detailPrice"`
	DetailQuantity string `xml:"detailQuantity"`
	GMOExtend2     string `xml:"gmoExtend2"`
	GMOExtend3     string `xml:"gmoExtend3"`
	GMOExtend4     string `xml:"gmoExtend4"`
	DetailBrand    string `xml:"detailBrand"`
	DetailCategory string `xml:"detailCategory"`
}

type Details []*Detail

func (c *Client) Register() {}
