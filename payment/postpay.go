package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// PostpayEntryTranRequest ... postpay entry tran request
type PostpayEntryTranRequest struct {
	OrderID string `schema:"OrderID" validate:"required,max=27"`
	Amount  int    `schema:"Amount" validate:"required"`
	Tax     int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *PostpayEntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// PostpayEntryTranResponse ... postpay entry tran response
type PostpayEntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// PostpayEntryTran ... postpay entry tran
func (cli *Client) PostpayEntryTran(
	req *PostpayEntryTranRequest,
) (*PostpayEntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PostpayEntryTranResponse{}
	_, err := cli.do(postpayEntryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PostpayExecTranRequest ... postpay exec tran request
type PostpayExecTranRequest struct {
	AccessID                 string                     `schema:"AccessID" validate:"required"`
	AccessPass               string                     `schema:"AccessPass" validate:"required"`
	OrderID                  string                     `schema:"OrderID" validate:"required"`
	HTTPHeaderAccept         string                     `schema:"HttpHeaderAccept,omitempty" validate:"omitempty"`
	HTTPHeaderAcceptCharset  string                     `schema:"HttpHeaderAcceptCharset,omitempty" validate:"omitempty"`
	HTTPHeaderAcceptEncoding string                     `schema:"HttpHeaderAcceptEncoding,omitempty" validate:"omitempty"`
	HTTPHeaderAcceptLanguage string                     `schema:"HttpHeaderAcceptLanguage,omitempty" validate:"omitempty"`
	HTTPHeaderClientIp       string                     `schema:"HttpHeaderClientIp,omitempty" validate:"omitempty"`
	HTTPHeaderConnection     string                     `schema:"HttpHeaderConnection,omitempty" validate:"omitempty"`
	HTTPHeaderDoNotTrack     string                     `schema:"HttpHeaderDoNotTrack,omitempty" validate:"omitempty"`
	HTTPHeaderHost           string                     `schema:"HttpHeaderHost,omitempty" validate:"omitempty"`
	HTTPHeaderReferrer       string                     `schema:"HttpHeaderReferrer,omitempty" validate:"omitempty"`
	HTTPHeaderUserAgent      string                     `schema:"HttpHeaderUserAgent,omitempty" validate:"omitempty"`
	HTTPHeaderKeepAlive      string                     `schema:"HttpHeaderKeepAlive,omitempty" validate:"omitempty"`
	HTTPHeaderUaCpu          string                     `schema:"HttpHeaderUaCpu,omitempty" validate:"omitempty"`
	HTTPHeaderVia            string                     `schema:"HttpHeaderVia,omitempty" validate:"omitempty"`
	HTTPHeaderXForwardedFor  string                     `schema:"HttpHeaderXForwardedFor,omitempty" validate:"omitempty"`
	HTTPHeaderOther          string                     `schema:"HttpHeaderOther,omitempty" validate:"omitempty"`
	CustomerIp               string                     `schema:"CustomerIp,omitempty" validate:"omitempty"`
	IMEI                     string                     `schema:"IMEI,omitempty" validate:"omitempty"`
	DeviceInfo               string                     `schema:"DeviceInfo,omitempty" validate:"omitempty"`
	CustomerOrderDate        string                     `schema:"CustomerOrderDate,omitempty" validate:"required,len=8"` //YYYYMMDD
	CustomerName             string                     `schema:"CustomerName,omitempty" validate:"required"`            // kanzi
	CustomerNameKana         string                     `schema:"CustomerNameKana,omitempty" validate:"required"`        // kana
	CustomerZipCode          string                     `schema:"CustomerZipCode,omitempty" validate:"required,len=7"`
	CustomerAddress          string                     `schema:"CustomerAddress,omitempty" validate:"required"`
	CustomerCompanyName      string                     `schema:"CustomerCompanyName,omitempty" validate:"omitempty"`
	CustomerDepartmentName   string                     `schema:"CustomerDepartmentName,omitempty" validate:"omitempty"`
	CustomerTel1             string                     `schema:"CustomerTel1,omitempty" validate:"required"`
	CustomerTel2             string                     `schema:"CustomerTel2,omitempty" validate:"omitempty"`
	CustomerEmail1           string                     `schema:"CustomerEmail1,omitempty" validate:"required,email"`
	CustomerEmail2           string                     `schema:"CustomerEmail2,omitempty" validate:"omitempty,email"`
	CustomerPaymentType      PostpayCustomerPaymentType `schema:"CustomerPaymentType,omitempty" validate:"required,numeric"`
	CustomerSex              PostpayCustomerSex         `schema:"CustomerSex,omitempty" validate:"omitempty,numeric"`
	CustomerBirthday         string                     `schema:"CustomerBirthday,omitempty" validate:"omitempty,len=8"`
	CustomerRegistDate       string                     `schema:"CustomerRegistDate,omitempty" validate:"omitempty,len=8"`
	CustomerBuyCount         int                        `schema:"CustomerBuyCount,omitempty" validate:"omitempty,numeric"`
	CustomerBuyAmountTotal   int                        `schema:"CustomerBuyAmountTotal,omitempty" validate:"omitempty,numeric"`
	CustomerID               string                     `schema:"CustomerID,omitempty" validate:"omitempty,alphanum"`
	DeliveryName             string                     `schema:"DeliveryName,omitempty" validate:"omitempty"`
	DeliveryNameKana         string                     `schema:"DeliveryNameKana,omitempty" validate:"omitempty"`
	DeliveryZipCode          string                     `schema:"DeliveryZipCode,omitempty" validate:"omitempty,len=7"`
	DeliveryAddress          string                     `schema:"DeliveryAddress,omitempty" validate:"omitempty"`
	DeliveryCompanyName      string                     `schema:"DeliveryCompanyName,omitempty" validate:"omitempty"`
	DeliveryDepartmentName   string                     `schema:"DeliveryDepartmentName,omitempty" validate:"omitempty"`
	DeliveryTel              string                     `schema:"DeliveryTel,omitempty" validate:"omitempty"`
	DetailName               string                     `schema:"DetailName,omitempty" validate:"omitempty"`
	DetailPrice              int                        `schema:"DetailPrice,omitempty" validate:"omitempty,numeric"`
	DetailQuantity           int                        `schema:"DetailQuantity,omitempty" validate:"omitempty,numeric"`
	DetailBrand              string                     `schema:"DetailBrand,omitempty" validate:"omitempty"`
	DetailCategory           string                     `schema:"DetailCategory,omitempty" validate:"omitempty"`
	ClientField1             string                     `schema:"ClientField1,omitempty" validate:"omitempty"`
	ClientField2             string                     `schema:"ClientField2,omitempty" validate:"omitempty"`
	ClientField3             string                     `schema:"ClientField3,omitempty" validate:"omitempty"`
}

// Validate ... validate
func (r *PostpayExecTranRequest) Validate() error {
	return validate.Struct(r)
}

// PostpayExecTranResponse ... postpay exec tran response
type PostpayExecTranResponse struct {
	OrderID              string           `schema:"OrderID,omitempty"`
	Status               TradeMultiStatus `schema:"Status,omitempty"`
	TranDate             string           `schema:"TranDate,omitempty"` //yyyyMMddHHmmss
	PostpayTransactionID string           `schema:"PostpayTransactionId,omitempty"`
	CheckString          string           `schema:"CheckString,omitempty"`
	ErrCode              string           `schema:"ErrCode,omitempty"`
	ErrInfo              string           `schema:"ErrInfo,omitempty"`
}

// PostpayExecTran ... postpay exec tran
func (cli *Client) PostpayExecTran(
	req *PostpayExecTranRequest,
) (*PostpayExecTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PostpayExecTranResponse{}
	_, err := cli.do(postpayExecTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PostpayShippedTranRequest struct {
	AccessID      string               `schema:"AccessID" validate:"required"`
	AccessPass    string               `schema:"AccessPass" validate:"required"`
	OrderID       string               `schema:"OrderID" validate:"required"`
	PDCompanyCode PostpayPDCompanyCode `schema:"PdCompanyCode" validate:"required"`
	SlipNo        string               `schema:"SlipNo" validate:"required"`
}

func (r *PostpayShippedTranRequest) Validate() error {
	return validate.Struct(r)
}

type PostpayShippedTranResponse struct {
	OrderID string           `schema:"OrderID,omitempty"`
	Status  TradeMultiStatus `schema:"Status,omitempty"`
	ErrCode string           `schema:"ErrCode,omitempty"`
	ErrInfo string           `schema:"ErrInfo,omitempty"`
}

func (cli *Client) PostpayShippedTran(
	req *PostpayShippedTranRequest,
) (*PostpayShippedTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &PostpayShippedTranResponse{}
	_, err := cli.do(postpayShippedTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
