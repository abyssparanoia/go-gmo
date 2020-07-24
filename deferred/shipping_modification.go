package deferred

import "context"

type shippingModifyRequest struct {
	ShopInfo    *shopInfo                  `xml:"shopInfo"`
	Transaction *shippingReportTransaction `xml:"transaction"`
	KindInfo    *kindInfo                  `xml:"kindInfo"`
}

type kindInfo struct {
	UpdateKind uint8 `xml:"updateKind"`
}

type shippingModifyResponse struct {
	Result            string             `xml:"result"`
	Errors            *gmoErrors         `xml:"errors"`
	TransactionResult *transactionResult `xml:"transactionResult"`
}

func (c *Client) ModifyShippingReport(ctx context.Context, req *ShippingModifyRequest) (*ShippingModifyResponse, error) {
	if req == nil {
		return nil, errInvalidParameterPassed
	}
	body := req.toParam()
	respParam := shippingModifyResponse{}
	body.ShopInfo = &shopInfo{
		AuthenticationID: c.AuthenticationID,
		ShopCode:         c.ShopCode,
		ConnectPassword:  c.ConnectPassword,
	}
	status, err := c.doAndUnmarshalXML(ctx, POST, c.APIHost, []string{"auto", "pdrequest.do"}, map[string]string{},
		body, &respParam)
	if err != nil {
		return nil, err
	}
	resp := newShippingModifyResponse(&respParam)
	resp.Status = status
	return resp, nil
}
