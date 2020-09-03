package deferred

import "context"

type transaction struct {
	GMOTransactionID string `xml:"gmoTransactionId" validate:"required"`
}

type authResultGetRequest struct {
	ShopInfo    *shopInfo    `xml:"shopInfo"`
	Transaction *transaction `xml:"transaction" validate:"required"`
}

type authResultGetResponse struct {
	Result            string             `xml:"result"`
	Errors            *gmoErrors         `xml:"errors"`
	TransactionResult *transactionResult `xml:"transactionResult"`
}

func (c *Client) GetAuthResult(ctx context.Context, req *AuthResultGetRequest) (*AuthResultGetResponse, error) {
	if req == nil {
		return nil, errInvalidParameterPassed
	}
	body, err := req.toParam()
	if err != nil {
		return nil, err
	}
	respParam := authResultGetResponse{}
	body.ShopInfo = &shopInfo{
		AuthenticationID: c.AuthenticationID,
		ShopCode:         c.ShopCode,
		ConnectPassword:  c.ConnectPassword,
	}
	status, err := c.doAndUnmarshalXML(ctx, POST, c.APIHost, []string{"auto", "getauthor.do"}, map[string]string{},
		body, &respParam)
	if err != nil {
		return nil, err
	}
	resp := newAuthResultGetResponse(&respParam)
	resp.Status = status
	return resp, nil
}
