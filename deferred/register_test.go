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

func TestDeferred_RegisterTransaction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *RegisterRequestParam
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
		want          *RegisterResponseParam
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &RegisterRequestParam{
					Buyer: &Buyer{
						ShopTransactionID: "1111-2222-3333-4444",
						ShopOrderDate:     "2012/04/01",
						FullName:          "山田太郎",
						FullNameKana:      "ヤマダタロウ",
						ZipCode:           "150-0043",
						Address:           "東京都渋谷区道玄坂1−14−6",
						CompanyName:       "",
						DepartmentName:    "",
						Tel1:              "0120-1111-2222",
						Tel2:              "0120-1111-2222",
						Email:             "gmo@gmo.biz",
						Email2:            "gmo@docomo.ne.jp",
						BilledAmount:      "10001",
						GMOExtend1:        "",
						PaymentType:       "2",
						Sex:               "1",
						BirthDay:          "1996/01/01",
						MemberRegistDate:  "2015/12/22",
						BuyCount:          "5",
						BuyAmountTotal:    "15000",
						MemberID:          "1234567890",
					},
					Deliveries: Deliveries{
						{
							DeliveryCustomer: &DeliveryCustomer{
								FullName:       "山田 太郎",
								FullNameKana:   "",
								ZipCode:        "150-0043",
								Address:        "東京都渋谷区道玄坂 1−14−6",
								CompanyName:    "",
								DepartmentName: "",
								Tel:            "0120-1111-2222",
							},
							Details: Details{
								{
									DetailName:     "鉛筆",
									DetailPrice:    "120",
									DetailQuantity: "1",
									GMOExtend2:     "",
									GMOExtend3:     "",
									GMOExtend4:     "",
									DetailBrand:    "鉛筆メーカー",
									DetailCategory: "文具",
								},
								{
									DetailName:     "消しゴム",
									DetailPrice:    "105",
									DetailQuantity: "2",
									GMOExtend2:     "",
									GMOExtend3:     "",
									GMOExtend4:     "",
									DetailBrand:    "消しゴムメーカー",
									DetailCategory: "文具",
								},
							},
						},
					},
				},
			},
			responseParam: responseParam{
				statusCode: http.StatusOK,
				body: `
				<?xml version="1.0" encoding="UTF-8"?>
				<response>
				<result>OK</result>
				<errors/>
				<transactionResult>
				<shopTransactionId>1111-2222-3333-4444</shopTransactionId>
				<gmoTransactionId>12900000018</gmoTransactionId>
				<authorResult>審査中</authorResult>
				</transactionResult>
				</response>
					`,
				header: map[string]string{},
			},
			want: &RegisterResponseParam{
				Result: "OK",
				Errors: Errors{},
				TransactionResult: &TransactionResult{
					ShopTransactionID: "1111-2222-3333-4444",
					GMOTransactionID:  "12900000018",
					AuthorResult:      "審査中",
					AutoAutherResult:  "",
					MaulAuthorResult:  "",
				},
				Status: 200,
			},
		},
		{
			name: "cancelled context",
			args: args{
				ctx: testutil.NewCanceledContext(),
				req: &RegisterRequestParam{
					Buyer: &Buyer{
						ShopTransactionID: "1111-2222-3333-4444",
						ShopOrderDate:     "2012/04/01",
						FullName:          "山田太郎",
						FullNameKana:      "ヤマダタロウ",
						ZipCode:           "150-0043",
						Address:           "東京都渋谷区道玄坂1−14−6",
						CompanyName:       "",
						DepartmentName:    "",
						Tel1:              "0120-1111-2222",
						Tel2:              "0120-1111-2222",
						Email:             "gmo@gmo.biz",
						Email2:            "gmo@docomo.ne.jp",
						BilledAmount:      "10001",
						GMOExtend1:        "",
						PaymentType:       "2",
						Sex:               "1",
						BirthDay:          "1996/01/01",
						MemberRegistDate:  "2015/12/22",
						BuyCount:          "5",
						BuyAmountTotal:    "15000",
						MemberID:          "1234567890",
					},
					Deliveries: Deliveries{
						{
							DeliveryCustomer: &DeliveryCustomer{
								FullName:       "山田 太郎",
								FullNameKana:   "",
								ZipCode:        "150-0043",
								Address:        "東京都渋谷区道玄坂 1−14−6",
								CompanyName:    "",
								DepartmentName: "",
								Tel:            "0120-1111-2222",
							},
							Details: Details{
								{
									DetailName:     "鉛筆",
									DetailPrice:    "120",
									DetailQuantity: "1",
									GMOExtend2:     "",
									GMOExtend3:     "",
									GMOExtend4:     "",
									DetailBrand:    "鉛筆メーカー",
									DetailCategory: "文具",
								},
								{
									DetailName:     "消しゴム",
									DetailPrice:    "105",
									DetailQuantity: "2",
									GMOExtend2:     "",
									GMOExtend3:     "",
									GMOExtend4:     "",
									DetailBrand:    "消しゴムメーカー",
									DetailCategory: "文具",
								},
							},
						},
					},
				},
			},
			responseParam: responseParam{},
			wantErr:       true,
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
			got, err := c.RegisterTransaction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("c.RegisterTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
