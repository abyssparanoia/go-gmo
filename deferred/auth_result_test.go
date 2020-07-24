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

func TestDeferred_GetAuthResult(t *testing.T) {
	type args struct {
		ctx context.Context
		req *AuthResultGetRequest
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
		want          *AuthResultGetResponse
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &AuthResultGetRequest{
					Transaction: &Transaction{
						GMOTransactionID: "1111-2222-3333-4444",
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
			want: &AuthResultGetResponse{
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
				req: &AuthResultGetRequest{
					Transaction: &Transaction{
						GMOTransactionID: "1111-2222-3333-4444",
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
			got, err := c.GetAuthResult(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("c.GetAuthResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
