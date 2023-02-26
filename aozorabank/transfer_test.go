package aozorabank

import (
	"context"
	"encoding/json"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/bxcodec/faker"
)

func TestGetTransferStatus(
	t *testing.T,
) {
	testcases := map[string]struct {
		request      *GetTransferStatusRequest
		expectedPath string
		expected     *GetTransferStatusResponse
	}{
		"ok": {
			request: &GetTransferStatusRequest{
				AccountID:               "111111111111",
				QueryKeyClass:           QueryKeyClassTransferApplies,
				ApplyNo:                 "2018072902345678",
				DateFrom:                "2018-07-30",
				DateTo:                  "2018-08-10",
				NextItemKey:             "1234567890",
				RequestTransferStatuses: []*RequestTransferStatus{{TransferStatusApplying}},
				RequestTransferClass:    RequestTransferClassAll,
				RequestTransferTerm:     RequestTransferTermTransferDesignatedDate,
			},
			expectedPath: "",
			expected:     fakeData(GetTransferStatusResponse{}),
		},
		"ok (required only)": {
			request: &GetTransferStatusRequest{
				AccountID:     "111111111111",
				QueryKeyClass: QueryKeyClassTransferApplies,
			},
			expected: fakeData(GetTransferStatusResponse{}),
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()

			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewClient(false, "testAccessToken")
			cli.APIHost = apiHostTest
			result, err := cli.GetTransferStatus(context.TODO(), tc.request)
			assert.Equal(t, nil, err)
			assert.Equal(t, expected, result)
		})
	}
}

func TestTransferRequest(
	t *testing.T,
) {
	testcases := map[string]struct {
		request  *TransferRequestRequest
		expected *TransferRequestResponse
	}{
		"ok": {
			request: &TransferRequestRequest{
				IdempotencyKey:          "111111111111",
				AccountID:               "101011234567",
				RemitterName:            "ｼﾞ-ｴﾑｵ-ｼｮｳｼﾞ(ｶ",
				TransferDesignatedDate:  "2018-07-30",
				TransferDateHolidayCode: TransferDateHolidayCodeNextBusinessDay,
				TotalCount:              0,
				TotalAmount:             1000,
				ApplyComment:            "緊急で承認をお願いします",
				Transfers: []*Transfer{
					{
						ItemID:                "1",
						TransferAmount:        100,
						EDIInfo:               "ｾｲｷﾕｳｼﾖﾊﾞﾝｺﾞｳ1234",
						BeneficiaryBankCode:   "0398",
						BeneficiaryBankName:   "ｱｵｿﾞﾗ",
						BeneficiaryBranchCode: "111",
						BeneficiaryBranchName: "ﾎﾝﾃﾝ",
						AccountTypeCode:       AccountTypeCodeOrdinary,
						AccountNumber:         "1234567",
						BeneficiaryName:       "ｶ)ｱｵｿﾞﾗｻﾝｷﾞｮｳ",
					},
				},
			},
			expected: fakeData(TransferRequestResponse{}),
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()

			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewClient(false, "testAccessToken")
			cli.APIHost = apiHostTest
			result, err := cli.TransferRequest(context.TODO(), tc.request)
			assert.Equal(t, nil, err)
			assert.Equal(t, expected, result)
		})
	}
}

func TestGetRequestResult(
	t *testing.T,
) {
	testcases := map[string]struct {
		request  *GetRequestResultRequest
		expected *GetRequestResultResponse
	}{
		"ok": {
			request: &GetRequestResultRequest{
				AccountID: "111111111111",
				ApplyNo:   "2018072902345678",
			},
			expected: fakeData(GetRequestResultResponse{}),
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()

			expected := tc.expected
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respBody, _ := json.Marshal(expected)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respBody)
			}))
			defer ts.Close()
			defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
			http.DefaultTransport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(ts.URL)
			}
			defer func() { http.DefaultTransport.(*http.Transport).Proxy = defaultProxy }()

			cli, _ := NewClient(false, "testAccessToken")
			cli.APIHost = apiHostTest
			result, err := cli.GetRequestResult(context.TODO(), tc.request)
			assert.Equal(t, nil, err)
			assert.Equal(t, expected, result)
		})
	}
}

func fakeData[T any](t T) *T {
	ret := new(T)
	if err := faker.FakeData(ret); err != nil {
		panic(err)
	}
	return ret
}
