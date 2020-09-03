package deferred

import "context"

type modifyRequest struct {
	ShopInfo   *shopInfo          `xml:"shopInfo"`
	Buyer      *buyerModification `xml:"buyer" validate:"required"`
	Deliveries deliveries         `xml:"deliveries"`
	KindInfo   *kindInfo          `xml:"kindInfo" validate:"required"`
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
