package deferred

import "context"

type invoiceRequest struct {
	ShopInfo    *shopInfo    `xml:"shopInfo"`
	Transaction *transaction `xml:"transaction" validate:"required"`
}

type invoiceResponse struct {
	GMOTransactionID              string `xml:"gmoTransactionID"`
	DeliveryZip                   string `xml:"deliveryZip"`
	DeliveryAddress1              string `xml:"deliveryAddress1"`
	DeliveryAddress2              string `xml:"deliveryAddress2"`
	PurchaseCompanyName           string `xml:"purchaseCompanyName"`
	PurchaseDepartmentName        string `xml:"purchaseDepartmentName"`
	PurchaseUserName              string `xml:"purchaseUserName"`
	ShopName                      string `xml:"shopName"`
	ShopTransactionID             string `xml:"shopTransactionID"`
	InvoiceMatter1                string `xml:"invoiceMatter1"`
	InvoiceMatter2                string `xml:"invoiceMatter2"`
	InvoiceMatter3                string `xml:"invoiceMatter3"`
	InvoiceMatter4                string `xml:"invoiceMatter4"`
	InvoiceMatter5                string `xml:"invoiceMatter5"`
	GMOCompanyName                string `xml:"gMOCompanyName"`
	GMOInfo1                      string `xml:"gMOInfo1"`
	GMOInfo2                      string `xml:"gMOInfo2"`
	GMOInfo3                      string `xml:"gMOInfo3"`
	GMOInfo4                      string `xml:"gMOInfo4"`
	InvoiceTitle                  string `xml:"invoiceTitle"`
	InvoiceGreeting1              string `xml:"invoiceGreeting1"`
	InvoiceGreeting2              string `xml:"invoiceGreeting2"`
	InvoiceGreeting3              string `xml:"invoiceGreeting3"`
	InvoiceGreeting4              string `xml:"invoiceGreeting4"`
	Yobi1                         string `xml:"yobi1"`
	Yobi2                         string `xml:"yobi2"`
	Yobi3                         string `xml:"yobi3"`
	Yobi4                         string `xml:"yobi4"`
	Yobi5                         string `xml:"yobi5"`
	Yobi6                         string `xml:"yobi6"`
	Yobi7                         string `xml:"yobi7"`
	Yobi8                         string `xml:"yobi8"`
	Yobi9                         string `xml:"yobi9"`
	Yobi10                        string `xml:"yobi10"`
	BilledAmount                  string `xml:"billedAmount"`
	BilledAmountTax               string `xml:"billedAmountTax"`
	OrderDate                     string `xml:"orderDate"`
	InvoiceIssueDate              string `xml:"invoiceIssueDate"`
	PaymentDueDate                string `xml:"paymentDueDate"`
	TrackingNumber                string `xml:"trackingNumber"`
	BankNoteWording               string `xml:"bankNoteWording"`
	BankName                      string `xml:"bankName"`
	Bankcode                      string `xml:"bankcode"`
	DepositType                   string `xml:"depositType"`
	AccountNumber                 string `xml:"accountNumber"`
	BankAccountHolder             string `xml:"bankAccountHolder"`
	VotesBilledAmount             string `xml:"votesBilledAmount"`
	VotesFrontUpperInfo           string `xml:"votesFrontUpperInfo"`
	VotesBarCode                  string `xml:"votesBarCode"`
	DocketBilledAmount            string `xml:"docketBilledAmount"`
	DocketPurchaseAddress         string `xml:"docketPurchaseAddress"`
	DocketPurchaseCompanyName     string `xml:"docketPurchaseCompanyName"`
	DocketPurchaseDepartmentName  string `xml:"docketPurchaseDepartmentName"`
	DocketPurchaseUserName        string `xml:"docketPurchaseUserName"`
	DocketTrackingNumber          string `xml:"docketTrackingNumber"`
	DocketX                       string `xml:"docketX"`
	ReceiptPurchaseCompanyName    string `xml:"receiptPurchaseCompanyName"`
	ReceiptPurchaseDepartmentName string `xml:"receiptPurchaseDepartmentName"`
	ReceiptPurchaseUserName       string `xml:"receiptPurchaseUserName"`
	ReceiptTrackingNumber1        string `xml:"receiptTrackingNumber1"`
	ReceiptTrackingNumber2        string `xml:"receiptTrackingNumber2"`
	ReceiptAmount                 string `xml:"receiptAmount"`
	ReceiptTax                    string `xml:"receiptTax"`
	ReceiptShopName               string `xml:"receiptShopName"`
	ReceiptPrintWord              string `xml:"receiptPrintWord"`
	String                        string `xml:"string"`
	Yobi11                        string `xml:"yobi11"`
	Yobi12                        string `xml:"yobi12"`
	Yobi13                        string `xml:"yobi13"`
	Yobi14                        string `xml:"yobi14"`
	Yobi15                        string `xml:"yobi15"`
	DetailList                    struct {
		GoodsDetail struct {
			GoodsName   string  `xml:"goodsName"`
			GoodsNum    int64   `xml:"goodsNum"`
			GoodsPrice  float64 `xml:"goodsPrice"`
			GoodsAmount float64 `xml:"goodsAmount"`
			GoodsTax    float64 `xml:"goodsTax"`
			Yobi16      string  `xml:"yobi16"`
			Yobi17      string  `xml:"yobi17"`
			Yobi18      string  `xml:"yobi18"`
			Yobi19      string  `xml:"yobi19"`
			Yobi20      string  `xml:"yobi20"`
		} `xml:"goodsDetail"`
	} `xml:"detailList"`
}

func (c *Client) GetInvoice(ctx context.Context, req *InvoiceGetRequest) (*InvoiceGetResponse, error) {
	if req == nil {
		return nil, errInvalidParameterPassed
	}
	body, err := req.toParam()
	if err != nil {
		return nil, err
	}
	respParam := invoiceResponse{}
	body.ShopInfo = &shopInfo{
		AuthenticationID: c.AuthenticationID,
		ShopCode:         c.ShopCode,
		ConnectPassword:  c.ConnectPassword,
	}
	status, err := c.doAndUnmarshalXML(ctx, POST, c.APIHost, []string{"auto", "getinvoicedata.do"}, map[string]string{},
		body, &respParam)
	if err != nil {
		return nil, err
	}
	resp := newInvoiceGetResponse(&respParam)
	resp.Status = status
	return resp, nil
}
