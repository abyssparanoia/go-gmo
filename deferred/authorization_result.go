package deferred

import "context"

type authorizationResultGetRequest struct {
	ShopInfo    *shopInfo    `xml:"shopInfo"`
	Transaction *transaction `xml:"transaction"`
}

type transaction struct {
	GMOTransactionID string `xml:"gmoTransactionId"`
}

type authorizationResultGetResponse struct {
	Result            string             `xml:"result"`
	Errors            Errors             `xml:"errors"`
	TransactionResult *TransactionResult `xml:"transactionResult"`
}

func (c *Client) GetAuthorizationResult(ctx context.Context, req *AuthorizationResultGetRequest) (*AuthorizationResultGetResponse, error) {
	body := req.toParam()
	respParam := authorizationResultGetResponse{}
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
	resp := newAuthorizationResultGetResponse(&respParam)
	resp.Status = status
	return resp, nil
}
