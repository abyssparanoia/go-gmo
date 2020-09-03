package deferred

import "context"

type modifyRequest struct {
	ShopInfo   *shopInfo          `xml:"shopInfo"`
	Buyer      *buyerModification `xml:"buyer" validate:"required"`
	Deliveries deliveries         `xml:"deliveries"`
	KindInfo   *kindInfo          `xml:"kindInfo" validate:"required"`
}

type buyerModification struct {
	GMOTransactionID  string `xml:"gmoTransactionId" validate:"required"`
	ShopTransactionID string `xml:"shopTransactionId,omitempty"`
	ShopOrderDate     string `xml:"shopOrderDate,omitempty"`
	FullName          string `xml:"fullName,omitempty"`
	FullNameKana      string `xml:"fullKanaName,omitempty"`
	ZipCode           string `xml:"zipCode,omitempty"`
	Address           string `xml:"address,omitempty"`
	CompanyName       string `xml:"companyName,omitempty"`
	DepartmentName    string `xml:"departmentName,omitempty"`
	Tel1              string `xml:"tel1,omitempty"`
	Tel2              string `xml:"tel2,omitempty"`
	Email             string `xml:"email1,omitempty"`
	Email2            string `xml:"email2,omitempty"`
	BilledAmount      string `xml:"billedAmount,omitempty"`
	GMOExtend1        string `xml:"gmoExtend1,omitempty"`
	PaymentType       string `xml:"paymentType,omitempty"`
	Sex               string `xml:"sex,omitempty"`
	BirthDay          string `xml:"birthDay,omitempty"`
	MemberRegistDate  string `xml:"memberRegistDate,omitempty"`
	BuyCount          string `xml:"buyCount,omitempty"`
	BuyAmountTotal    string `xml:"buyAmoutTotal,omitempty"`
	MemberID          string `xml:"memberId,omitempty"`
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

type modifyResponse struct {
	Result            string             `xml:"result"`
	Errors            *gmoErrors         `xml:"errors"`
	TransactionResult *transactionResult `xml:"transactionResult"`
}

func (c *Client) ModifyTransaction(ctx context.Context, req *ModifyRequest) (*ModifyResponse, error) {
	if req == nil {
		return nil, errInvalidParameterPassed
	}
	body, err := req.toParam()
	if err != nil {
		return nil, err
	}
	respParam := modifyResponse{}
	body.ShopInfo = &shopInfo{
		AuthenticationID: c.AuthenticationID,
		ShopCode:         c.ShopCode,
		ConnectPassword:  c.ConnectPassword,
	}
	status, err := c.doAndUnmarshalXML(ctx, POST, c.APIHost, []string{"auto", "modifycanceltransaction.do"}, map[string]string{},
		body, &respParam)
	if err != nil {
		return nil, err
	}
	resp := newModifyResponse(&respParam)
	resp.Status = status
	return resp, nil
}
