package deferred

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/abyssparanoia/go-gmo/deferred/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDeferred_GetInvoice(t *testing.T) {
	type args struct {
		ctx context.Context
		req *InvoiceGetRequest
	}
	type responseParam struct {
		body       string
		statusCode int
		header     map[string]string
	}
	tests := []struct {
		name          string
		args          args
		responseParam responseParam
		want          *InvoiceGetResponse
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &InvoiceGetRequest{
					GMOTransactionID: "20090300001",
				},
			},
			responseParam: responseParam{
				statusCode: http.StatusOK,
				body: `
				<?xml version="1.0" encoding="UTF-8"?>
				<response>
				<result>OK</result>
				<invoiceDataResult>
					<gmoTransactionId>123456789012</gmoTransactionId>
					<deliveryZip>1500043</deliveryZip>
					<deliveryAddress1>東京都渋谷区道玄坂1-1-1</deliveryAddress1>
					<deliveryAddress2></deliveryAddress2>
					<purchaseCompanyName></purchaseCompanyName>
					<purchaseDepartmentName></purchaseDepartmentName>
					<purchaseUserName>山田太郎</purchaseUserName>
					<shopName>○○サービス</shopName>
					<shopTransactionId>1111-2222-3333-4444</shopTransactionId>
					<invoiceMatter1></invoiceMatter1>
					<invoiceMatter2></invoiceMatter2>
					<invoiceMatter3></invoiceMatter3>
					<invoiceMatter4></invoiceMatter4>
					<invoiceMatter5></invoiceMatter5>
					<gmoCompanyName>GMO ペイメントサービス</gmoCompanyName>
					<gmoInfo1>http://www.gmo-ps.co.jp</gmoInfo1>
					<gmoInfo2>03-xxxx-xxxx</gmoInfo2>
					<gmoInfo3>xxx@xxx.xxx</gmoInfo3>
					<gmoInfo4></gmoInfo4>
					<invoiceTitle>○○サービス</invoiceTitle>
					<invoiceGreeting1>このたびは</invoiceGreeting1>
					<invoiceGreeting2>お買い上げ頂きまして</invoiceGreeting2>
					<invoiceGreeting3>誠に</invoiceGreeting3>
					<invoiceGreeting4>ありがとうございます。</invoiceGreeting4>
					<yobi1></yobi1>
					<yobi2></yobi2>
					<yobi3></yobi3>
					<yobi4></yobi4>
					<yobi5></yobi5>
					<yobi6></yobi6>
					<yobi7></yobi7>
					<yobi8></yobi8>
					<yobi9></yobi9>
					<yobi10></yobi10>
					<billedAmount>10500</billedAmount>
					<billedAmountTax>500</billedAmountTax>
					<orderDate>2013/05/01</orderDate>
					<invoiceIssueDate>2013/05/07</invoiceIssueDate>
					<paymentDueDate>2013/05/21</paymentDueDate>
					<trackingNumber>ab000000-00123456789012</trackingNumber>
					<bankNoteWording>振込先にご注意ください</bankNoteWording>
					<bankName>ジャパンネット</bankName>
					<bankCode>0000</bankCode>
					<branchName>サクラ</branchName>
					<branchCode>000</branchCode>
					<depositType>普通</depositType>
					<accountNumber>1234567</accountNumber>
					<bankAccountHolder>カ)ジーエムオー</bankAccountHolder>
					<votesBilledAmount>10500</votesBilledAmount>
					<votesFontUpperInfo></votesFontUpperInfo>
					<votesFontKiwerInfo></votesFontKiwerInfo>
					<votesPaymentDueDate>2013/05/21</votesPaymentDueDate>
					<votesPurchaseUserName>山田太郎</votesPurchaseUserName>
					<votesTrackingNumber>ab000000-00123456789012</votesTrackingNumber>
					<votesBarCode>12345678901234567890123456789012345678901234</votesBarCode>
					<docketBilledAmount>10500</docketBilledAmount>
					<docketPurchaseAddress>東京都千代田区</docketPurchaseAddress>
					<docketPurchaseCompanyName></docketPurchaseCompanyName>
					<docketPurchaseDepartmentName></docketPurchaseDepartmentName>
					<docketPurchaseUserName>山田太郎</docketPurchaseUserName>
					<docketTrackingNumber>ab000000-00123456789012</docketTrackingNumber>
					<docketX>X</docketX>
					<receiptPurchaseCompanyName></receiptPurchaseCompanyName>
					<receiptPurchaseDepartmentName></receiptPurchaseDepartmentName>
					<receiptPurchaseUserName>山田太郎</receiptPurchaseUserName>
					<receiptTrackingNumber1>ab000000-00123456789012</receiptTrackingNumber1>
					<receiptTrackingNumber2>ab000000-00123456789012</receiptTrackingNumber2>
					<receiptAmount>10500</receiptAmount>
					<receiptTax>500</receiptTax>
					<receiptShopName>○○サービス</receiptShopName>
					<receiptPrintWord>△△△</receiptPrintWord>
					<string></string>
					<yobi11></yobi11>
					<yobi12></yobi12>
					<yobi13></yobi13>
					<yobi14></yobi14>
					<yobi15></yobi15>
					<detailList>
						<goodsDetail>
							<goodsName>鉛筆</goodsName>
							<goodsNum>50</goodsNum>
							<goodsPrice>105</goodsPrice>
							<goodsAmount>5250</goodsAmount>
							<goodsAmountTax></goodsAmountTax>
							<yobi16></yobi16>
							<yobi17></yobi17>
							<yobi18></yobi18>
							<yobi19></yobi19>
							<yobi20></yobi20>
						</goodsDetail>
						<goodsDetail>
							<goodsName>消しゴム</goodsName>
							<goodsNum>50</goodsNum>
							<goodsPrice>105</goodsPrice>
							<goodsAmount>5250</goodsAmount>
							<goodsAmountTax></goodsAmountTax>
							<yobi16></yobi16>
							<yobi17></yobi17>
							<yobi18></yobi18>
							<yobi19></yobi19>
							<yobi20></yobi20>
						</goodsDetail>
					</detailList>
				</invoiceDataResult>
				<errors />
			</response>
					`,
				header: map[string]string{},
			},
			want: &InvoiceGetResponse{
				Status:                  200,
				DeliveryZip:             "1500043",
				DeliveryAddress1:        "東京都渋谷区道玄坂1-1-1",
				PurchaseUserName:        "山田太郎",
				ShopName:                "○○サービス",
				InvoiceTitle:            "○○サービス",
				InvoiceGreeting1:        "このたびは",
				InvoiceGreeting2:        "お買い上げ頂きまして",
				InvoiceGreeting3:        "誠に",
				InvoiceGreeting4:        "ありがとうございます。",
				BilledAmount:            "10500",
				BilledAmountTax:         "500",
				OrderDate:               "2013/05/01",
				InvoiceIssueDate:        "2013/05/07",
				PaymentDueDate:          "2013/05/21",
				TrackingNumber:          "ab000000-00123456789012",
				BankNoteWording:         "振込先にご注意ください",
				BankName:                "ジャパンネット",
				DepositType:             "普通",
				AccountNumber:           "1234567",
				BankAccountHolder:       "カ)ジーエムオー",
				VotesBilledAmount:       "10500",
				VotesBarCode:            "12345678901234567890123456789012345678901234",
				DocketBilledAmount:      "10500",
				DocketPurchaseAddress:   "東京都千代田区",
				DocketPurchaseUserName:  "山田太郎",
				DocketTrackingNumber:    "ab000000-00123456789012",
				DocketX:                 "X",
				ReceiptPurchaseUserName: "山田太郎",
				ReceiptTrackingNumber1:  "ab000000-00123456789012",
				ReceiptTrackingNumber2:  "ab000000-00123456789012",
				ReceiptAmount:           "10500",
				ReceiptTax:              "500",
				ReceiptShopName:         "○○サービス",
				ReceiptPrintWord:        "△△△",
				GoodsDetail: GoodsDetail{
					GoodsName:   "消しゴム",
					GoodsNum:    50,
					GoodsPrice:  105,
					GoodsAmount: 5250,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := testutil.NewTestClient(func(req *http.Request) *http.Response {
				header := make(http.Header)
				for k, v := range tt.responseParam.header {
					header.Set(k, v)
				}
				return &http.Response{
					StatusCode: tt.responseParam.statusCode,
					Body:       ioutil.NopCloser(bytes.NewBufferString(tt.responseParam.body)),
					Header:     header,
				}
			})
			defaultHTTPClient = client
			c, err := NewClient("XXX", "YYY", "ZZZ", true)
			got, err := c.GetInvoice(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("c.GetInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
