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

func TestDeferred_ModifyShippingReport(t *testing.T) {
	type args struct {
		ctx context.Context
		req *ShippingModifyRequest
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
		want          *ShippingModifyResponse
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &ShippingModifyRequest{
					Transaction: &ShippingReportTransaction{
						GMOTransactionID: "123456789012",
						PDCompanyCode:    "11",
						SlipNo:           "1234567890",
					},
					KindInfo: &KindInfo{
						UpdateKind: Cancel,
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
				<gmoTransactionId>12900000018</gmoTransactionId>
				</transactionResult>
				</response>
					`,
				header: map[string]string{},
			},
			want: &ShippingModifyResponse{
				Result: "OK",
				Errors: Errors{},
				TransactionResult: &TransactionResult{
					GMOTransactionID: "12900000018",
				},
				Status: 200,
			},
		},
		{
			name: "cancelled context",
			args: args{
				ctx: testutil.NewCanceledContext(),
				req: &ShippingModifyRequest{
					Transaction: &ShippingReportTransaction{
						GMOTransactionID: "123456789012",
						PDCompanyCode:    "11",
						SlipNo:           "1234567890",
					},
					KindInfo: &KindInfo{
						UpdateKind: Cancel,
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
			got, err := c.ModifyShippingReport(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("c.ModifyShippingReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
