package aozorabank

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

const (
	corporationPathV1 = "corporation/v1"
)

type (
	GetTransferStatusRequest struct {
		AccessToken             string                   `json:"-" validate:"required,min=1,max=128"`
		AccountID               string                   `json:"accountId" validate:"required,min=12,max=29"`
		QueryKeyClass           QueryKeyClass            `json:"queryKeyClass,string" validate:"required,oneof=1 2"`
		ApplyNo                 string                   `json:"applyNo,omitempty" validate:"omitempty,len=16"`
		DateFrom                string                   `json:"dateFrom,omitempty" validate:"omitempty,len=10"`
		DateTo                  string                   `json:"dateTo,omitempty" validate:"omitempty,len=10"`
		NextItemKey             string                   `json:"nextItemKey,omitempty" validate:"omitempty,min=1,max=24"`
		RequestTransferStatuses []*RequestTransferStatus `json:"requestTransferStatus,omitempty" validate:"omitempty"`
		RequestTransferClass    RequestTransferClass     `json:"requestTransferClass,string,omitempty" validate:"omitempty,oneof=1 3"`
		RequestTransferTerm     RequestTransferTerm      `json:"requestTransferTerm,string,omitempty" validate:"omitempty,oneof=1 2"`
	}

	RequestTransferStatus struct {
		Status TransferStatus `json:"requestTransferStatus,string"`
	}

	GetTransferStatusResponse struct {
		AcceptanceKeyClass         string                       `json:"acceptanceKeyClass"`
		BaseDate                   string                       `json:"baseDate"`
		BaseTime                   string                       `json:"baseTime"`
		Count                      int                          `json:"count,string"`
		TransferQueryBulkResponses []*transferQueryBulkResponse `json:"transferQueryBulkResponses"`
		TransferDetails            []*transferDetail            `json:"transferDetails"`
	}

	transferQueryBulkResponse struct {
		DateFrom                string                   `json:"dateFrom"`
		DateTo                  string                   `json:"dateTo"`
		RequestNextItemKey      string                   `json:"requestNextItemKey"`
		RequestTransferStatuses []*RequestTransferStatus `json:"requestTransferStatuses"`
		RequestTransferClass    string                   `json:"requestTransferClass"`
		RequestTransferTerm     string                   `json:"requestTransferTerm"`
		HasNext                 bool                     `json:"hasNext"`
		NextItemKey             string                   `json:"nextItemKey"`
	}

	transferDetail struct {
		TransferStatus     TransferStatus      `json:"transferStatus,string"`
		TransferStatusName string              `json:"transferStatusName"`
		TransferTypeName   string              `json:"transferTypeName"`
		IsFeeFreeUse       bool                `json:"isFeeFreeUse"`
		IsFeePointUse      bool                `json:"isFeePointUse"`
		PointName          string              `json:"pointName"`
		FeeLaterPaymentFlg bool                `json:"feeLaterPaymentFlg"`
		TransferDetailFee  string              `json:"transferDetailFee"`
		TotalDebitAmount   string              `json:"totalDebitAmount"`
		TransferApplies    []*transferApply    `json:"transferApplies"`
		TransferAccepts    []*transferAccept   `json:"transferAccepts"`
		TransferResponses  []*transferResponse `json:"transferResponses"`
	}

	transferApply struct {
		ApplyNo              string                 `json:"applyNo"`
		TransferApplyDetails []*transferApplyDetail `json:"transferApplyDetails"`
	}

	transferApplyDetail struct {
		ApplyDatetime   string `json:"applyDatetime"`
		ApplyStatus     string `json:"applyStatus"`
		ApplyUser       string `json:"applyUser"`
		ApplyComment    string `json:"applyComment"`
		ApprovalUser    string `json:"approvalUser"`
		ApprovalComment string `json:"approvalComment"`
	}

	transferAccept struct {
		AcceptNo       string `json:"acceptNo"`
		AcceptDatetime string `json:"acceptDatetime"`
	}

	transferResponse struct {
		AccountID              string          `json:"accountID"`
		RemitterName           string          `json:"remitterName"`
		TransferDesignatedDate string          `json:"transferDesignatedDate"`
		TransferInfos          []*transferInfo `json:"transferInfos"`
	}

	transferInfo struct {
		TransferAmount          string                    `json:"transferAmount"`
		EdiInfo                 string                    `json:"ediInfo"`
		BeneficiaryBankCode     string                    `json:"beneficiaryBankCode"`
		BeneficiaryBankName     string                    `json:"beneficiaryBankName"`
		BeneficiaryBranchCode   string                    `json:"beneficiaryBranchCode"`
		BeneficiaryBranchName   string                    `json:"beneficiaryBranchName"`
		AccountTypeCode         string                    `json:"accountTypeCode"`
		AccountNumber           string                    `json:"accountNumber"`
		BeneficiaryName         string                    `json:"beneficiaryName"`
		TransferDetailResponses []*transferDetailResponse `json:"transferDetailResponses"`
		UnableDetailInfos       []*unableDetailInfo       `json:"unableDetailInfos"`
	}

	transferDetailResponse struct {
		BeneficiaryBankNameKanji   string `json:"beneficiaryBankNameKanji"`
		BeneficiaryBranchNameKanji string `json:"beneficiaryBranchNameKanji"`
		UsedPoint                  string `json:"usedPoint"`
		IsFeeFreeUsed              bool   `json:"isFeeFreeUsed"`
		TransferFee                string `json:"applyNo"`
	}

	unableDetailInfo struct {
		TransferDetailStatus string `json:"transferDetailStatus"`
		RefundStatus         string `json:"refundStatus"`
		IsRepayment          bool   `json:"isRepayment"`
		RepaymentDate        string `json:"repaymentDate"`
	}
)

func (r *GetTransferStatusRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) GetTransferStatus(
	ctx context.Context,
	req *GetTransferStatusRequest,
) (*GetTransferStatusResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &GetTransferStatusResponse{}
	if _, err := cli.doGet(getTransferHeader(req.AccessToken), fmt.Sprintf("%s/transfer/status", corporationPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type (
	TransferRequestRequest struct {
		AccessToken             string                  `json:"-" validate:"required,min=1,max=128"`
		IdempotencyKey          string                  `json:"-" validate:"omitempty,min=1,max=128"`
		AccountID               string                  `json:"accountId" validate:"required,min=12,max=29"`
		RemitterName            string                  `json:"remitterName" validate:"omitempty,min=1,max=48"`
		TransferDesignatedDate  string                  `json:"transferDesignatedDate" validate:"omitempty"`
		TransferDateHolidayCode TransferDateHolidayCode `json:"transferDateHolidayCode,string" validate:"required,len=1"`
		TotalCount              int                     `json:"totalCount,string" validate:"omitempty,min=1,max=999999"`
		TotalAmount             int                     `json:"totalAmount,string" validate:"omitempty,min=1,max=999999999999"`
		ApplyComment            string                  `json:"applyComment" validate:"omitempty,min=1,max=20"`
		Transfers               []*Transfer             `json:"transfers" validate:"required"`
	}

	Transfer struct {
		ItemID                string          `json:"itemId" validate:"omitempty,min=1,max=6"`
		TransferAmount        int             `json:"transferAmount,string" validate:"required,min=1,max=20"`
		EDIInfo               string          `json:"ediInfo" validate:"omitempty,min=1,max=20"`
		BeneficiaryBankCode   string          `json:"beneficiaryBankCode" validate:"required,len=4"`
		BeneficiaryBankName   string          `json:"beneficiaryBankName" validate:"omitempty,min=1,max=30"`
		BeneficiaryBranchCode string          `json:"beneficiaryBranchCode" validate:"required,len=3"`
		BeneficiaryBranchName string          `json:"beneficiaryBranchName" validate:"omitempty,min=1,max=15"`
		AccountTypeCode       AccountTypeCode `json:"accountTypeCode,string" validate:"required,len=1"`
		AccountNumber         string          `json:"accountNumber" validate:"required,len=7"`
		BeneficiaryName       string          `json:"beneficiaryName" validate:"required,min=1,max=48"`
	}

	TransferRequestResponse struct {
		AccountID        string     `json:"accountId"`
		ResultCode       ResultCode `json:"resultCode,string"`
		ApplyNo          string     `json:"applyNo"`
		ApplyEndDatetime string     `json:"applyEndDatetime"`
	}
)

func (r *TransferRequestRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) TransferRequest(
	ctx context.Context,
	req *TransferRequestRequest,
) (*TransferRequestResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	header := getTransferHeader(req.AccessToken)
	header.Set(IdempotencyKeyHeaderKey, req.IdempotencyKey)
	res := &TransferRequestResponse{}
	if _, err := cli.doPost(header, fmt.Sprintf("%s/transfer/request", corporationPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type GetRequestResultRequest struct {
	AccessToken string `json:"-" validate:"required,min=1,max=128"`
	AccountID   string `json:"accountId" validate:"required,min=12,max=29"`
	ApplyNo     string `json:"applyNo" validate:"omitempty,len=16"`
}

type GetRequestResultResponse struct {
	AccountID        string     `json:"accountId"`
	ResultCode       ResultCode `json:"resultCode,string"`
	ApplyNo          string     `json:"applyNo"`
	ApplyEndDatetime string     `json:"applyEndDatetime"`
}

func (r *GetRequestResultRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) GetRequestResult(
	ctx context.Context,
	req *GetRequestResultRequest,
) (*GetRequestResultResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &GetRequestResultResponse{}
	if _, err := cli.doGet(getTransferHeader(req.AccessToken), fmt.Sprintf("%s/transfer/request-result", corporationPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type (
	GetBulkTransferStatusRequest struct {
		AccessToken             string                   `json:"-" validate:"required,min=1,max=128"`
		AccountID               string                   `json:"accountId" validate:"required,min=12,max=29"`
		QueryKeyClass           QueryKeyClass            `json:"queryKeyClass,string" validate:"required,oneof=1 2"`
		ApplyNo                 string                   `json:"applyNo,omitempty" validate:"omitempty,len=16"`
		DateFrom                string                   `json:"dateFrom,omitempty" validate:"omitempty,len=10"`
		DateTo                  string                   `json:"dateTo,omitempty" validate:"omitempty,len=10"`
		NextItemKey             string                   `json:"nextItemKey,omitempty" validate:"omitempty,min=1,max=24"`
		RequestTransferStatuses []*RequestTransferStatus `json:"requestTransferStatus,omitempty" validate:"omitempty"`
		RequestTransferClass    RequestTransferClass     `json:"requestTransferClass,string,omitempty" validate:"omitempty,oneof=1 3"`
		RequestTransferTerm     RequestTransferTerm      `json:"requestTransferTerm,string,omitempty" validate:"omitempty,oneof=1 2"`
	}

	GetBulkTransferStatusResponse struct {
		AcceptanceKeyClass         string                       `json:"acceptanceKeyClass"`
		BaseDate                   string                       `json:"baseDate"`
		BaseTime                   string                       `json:"baseTime"`
		Count                      int                          `json:"count,string"`
		TransferQueryBulkResponses []*transferQueryBulkResponse `json:"transferQueryBulkResponses"`
		BulkTransferDetails        []*bulkTransferDetail        `json:"bulkTransferDetails"`
	}

	bulkTransferDetail struct {
		TransferStatus     TransferStatus          `json:"transferStatus,string"`
		TransferStatusName string                  `json:"transferStatusName"`
		TransferTypeName   string                  `json:"transferTypeName"`
		IsFeeFreeUse       bool                    `json:"isFeeFreeUse"`
		IsFeePointUse      bool                    `json:"isFeePointUse"`
		PointName          string                  `json:"pointName"`
		FeeLaterPaymentFlg bool                    `json:"feeLaterPaymentFlg"`
		TransferDetailFee  string                  `json:"transferDetailFee"`
		TotalFee           string                  `json:"totalFee"`
		TotalDebitAmount   string                  `json:"totalDebitAmount"`
		TransferApplies    []*transferApply        `json:"transferApplies"`
		TransferAccepts    []*transferAccept       `json:"transferAccepts"`
		TransferResponses  []*bulkTransferResponse `json:"bulktransferResponses"`
	}

	bulkTransferResponse struct {
		AccountID              string          `json:"accountID"`
		RemitterName           string          `json:"remitterName"`
		TransferDesignatedDate string          `json:"transferDesignatedDate"`
		TotalCount             string          `json:"totalCount,string"`
		TotalAmount            string          `json:"totalAmount"`
		TransferInfos          []*transferInfo `json:"bulkTransferInfos"`
	}
)

func (r *GetBulkTransferStatusRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) GetBulkTransferStatus(
	ctx context.Context,
	req *GetBulkTransferStatusRequest,
) (*GetBulkTransferStatusResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	header := getTransferHeader(req.AccessToken)
	res := &GetBulkTransferStatusResponse{}
	if _, err := cli.doGet(header, fmt.Sprintf("%s/bulktransfer/status", corporationPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type (
	BulkTransferRequestRequest struct {
		AccessToken             string                  `json:"-" validate:"required,min=1,max=128"`
		IdempotencyKey          string                  `json:"-" validate:"omitempty,min=1,max=128"`
		AccountID               string                  `json:"accountId" validate:"required,min=12,max=29"`
		RemitterName            string                  `json:"remitterName" validate:"omitempty,min=1,max=48"`
		TransferDesignatedDate  string                  `json:"transferDesignatedDate" validate:"omitempty"`
		TransferDateHolidayCode TransferDateHolidayCode `json:"transferDateHolidayCode,string" validate:"required,len=1"`
		TotalCount              int                     `json:"totalCount,string" validate:"omitempty,min=1,max=999999"`
		TotalAmount             int                     `json:"totalAmount,string" validate:"omitempty,min=1,max=999999999999"`
		ApplyComment            string                  `json:"applyComment" validate:"omitempty,min=1,max=20"`
		BulkTransfers           []*BulkTransfer         `json:"bulkTransfers" validate:"required"`
	}

	BulkTransfer struct {
		ItemID                string          `json:"itemId" validate:"omitempty,min=1,max=6"`
		TransferAmount        int             `json:"transferAmount,string" validate:"required,min=1,max=20"`
		EDIInfo               string          `json:"ediInfo" validate:"omitempty,min=1,max=20"`
		BeneficiaryBankCode   string          `json:"beneficiaryBankCode" validate:"required,len=4"`
		BeneficiaryBankName   string          `json:"beneficiaryBankName" validate:"omitempty,min=1,max=30"`
		BeneficiaryBranchCode string          `json:"beneficiaryBranchCode" validate:"required,len=3"`
		BeneficiaryBranchName string          `json:"beneficiaryBranchName" validate:"omitempty,min=1,max=15"`
		AccountTypeCode       AccountTypeCode `json:"accountTypeCode,string" validate:"required,len=1"`
		AccountNumber         string          `json:"accountNumber" validate:"required,len=7"`
		BeneficiaryName       string          `json:"beneficiaryName" validate:"required,min=1,max=48"`
	}

	BulkTransferRequestResponse struct {
		AccountID        string     `json:"accountId"`
		ResultCode       ResultCode `json:"resultCode,string"`
		ApplyNo          string     `json:"applyNo"`
		ApplyEndDatetime string     `json:"applyEndDatetime"`
	}
)

func (r *BulkTransferRequestRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) BulkTransferRequest(
	ctx context.Context,
	req *BulkTransferRequestRequest,
) (*BulkTransferRequestResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	header := getTransferHeader(req.AccessToken)
	header.Set(IdempotencyKeyHeaderKey, req.IdempotencyKey)
	res := &BulkTransferRequestResponse{}
	if _, err := cli.doPost(header, fmt.Sprintf("%s/bulktransfer/request", corporationPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

type GetBulkTransferRequestResultRequest struct {
	AccessToken string `json:"-" validate:"required,min=1,max=128"`
	AccountID   string `json:"accountId" validate:"required,min=12,max=29"`
	ApplyNo     string `json:"applyNo" validate:"omitempty,len=16"`
}

type GetBulkTransferRequestResultResponse struct {
	AccountID        string     `json:"accountId"`
	ResultCode       ResultCode `json:"resultCode,string"`
	ApplyNo          string     `json:"applyNo"`
	ApplyEndDatetime string     `json:"applyEndDatetime"`
}

func (r *GetBulkTransferRequestResultRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) GetBulkTransferRequestResult(
	ctx context.Context,
	req *GetBulkTransferRequestResultRequest,
) (*GetBulkTransferRequestResultResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &GetBulkTransferRequestResultResponse{}
	if _, err := cli.doGet(getTransferHeader(req.AccessToken), fmt.Sprintf("%s/bulktransfer/request-result", corporationPathV1), reqMap, res); err != nil {
		return nil, err
	}
	return res, nil
}

func getTransferHeader(
	accessToken string,
) http.Header {
	return http.Header{
		"Content-Type":   []string{"application/json"},
		"x-access-token": []string{accessToken},
	}
}
