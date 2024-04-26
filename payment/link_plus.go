package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type GetLinkPlusURLRequest struct {
	GuideMailSendFlag       int
	ThanksMailSendFlag      int
	BccSendMailAddress      string
	FromSendMailAddress     string
	FromSendMailName        string
	SendMailAddress         string
	SendMailName            string
	TemplateNo              string
	GuideSmsSendFlag        int
	SMSTelno                int
	AuthenticationQuestion1 string
	AuthenticationAnswer1   string
	AuthenticationQuestion2 string
	AuthenticationAnswer2   string
	AuthenticationRetryMax  int

	ConfigID            string `validate:"required"`
	OrderID             string `validate:"required"`
	Amount              int    `validate:"required"`
	Tax                 int
	ClientField1        string
	ClientField2        string
	ClientField3        string
	Overview            string
	Detail              string
	PayMethods          []LinkPlusPayMethod
	RetURL              string
	CompleteURL         string
	CancelURL           string
	NotifyMailAddress   string
	RetryMax            int
	ExpireDays          int
	PaymentExpireDate   string
	ResultSkipFlag      int
	ConfirmSkipFlag     int
	TranDetailShowFlag  int
	CustomerMailAddress string
	CustomerName        string
	CustomerKana        string
	TelNo               string
	TemplateID          LinkPlusTemplateID
	LogoURL             string
	ShopName            string
	ShopURL             string
	ColorPattern        LinkPlusColorPattern
	Lang                LinkPlusLang
	ShopDomain          string

	CreditCard *GetLinkPlusURLRequestCreditCard
}

func (r *GetLinkPlusURLRequest) Validate() error {
	return validate.Struct(r)
}

type GetLinkPlusURLRequestCreditCard struct {
	JobCD               JobCD
	Method              int
	PayTimes            int
	ItemCode            string
	TdFlag              int
	TdTenantName        string
	MemberID            string
	SecCodeRequiredFlag int
	SecCodeHiddenFlag   int
	Tds2Type            int
	RegistMemberID      string
	CardInputHiddenFlag int
	CardMaxCnt          int
}

type GetLinkPlusURLResponse struct {
	OrderID     string
	LinkURL     string
	ProcessDate string
	WarnList    []*GetLinkPlusURLResponseWarnList
	ErrCode     string
	ErrInfo     string
}

type GetLinkPlusURLResponseWarnList struct {
	WarnCode string
	WarnInfo string
}

type getLinkPlusURLRequestJSON struct {
	GetURLParam    getLinkPlusURLRequestJSONGetURLParam    `json:"geturlparam"`
	ConfigID       string                                  `json:"configid"`
	Transaction    getLinkPlusURLRequestJSONTransaction    `json:"transaction"`
	Customer       getLinkPlusURLRequestJSONCustomer       `json:"customer,omitempty"`
	DisplaySetting getLinkPlusURLRequestJSONDisplaySetting `json:"displaysetting,omitempty"`
	CreditCard     getLinkPlusURLRequestJSONCreditCard     `json:"credit,omitempty"`
}

type getLinkPlusURLRequestJSONGetURLParam struct {
	ShopID                  string `json:"ShopID"`
	ShopPass                string `json:"ShopPass"`
	GuideMailSendFlag       int    `json:"GuideMailSendFlag,omitempty"`
	ThanksMailSendFlag      int    `json:"ThanksMailSendFlag,omitempty"`
	BccSendMailAddress      string `json:"BccSendMailAddress,omitempty"`
	FromSendMailAddress     string `json:"FromSendMailAddress,omitempty"`
	FromSendMailName        string `json:"FromSendMailName,omitempty"`
	SendMailAddress         string `json:"SendMailAddress,omitempty"`
	SendMailName            string `json:"SendMailName,omitempty"`
	CustomerName            string `json:"CustomerName,omitempty"`
	TemplateNo              string `json:"TemplateNo,omitempty"`
	GuideSmsSendFlag        int    `json:"GuideSmsSendFlag,omitempty"`
	SMSTelno                int    `json:"SMSTelno,omitempty"`
	AuthenticationQuestion1 string `json:"AuthenticationQuestion1,omitempty"`
	AuthenticationAnswer1   string `json:"AuthenticationAnswer1,omitempty"`
	AuthenticationQuestion2 string `json:"AuthenticationQuestion2,omitempty"`
	AuthenticationAnswer2   string `json:"AuthenticationAnswer2,omitempty"`
	AuthenticationRetryMax  int    `json:"AuthenticationRetryMax,omitempty"`
}

type getLinkPlusURLRequestJSONTransaction struct {
	OrderID            string              `json:"OrderID"`
	Amount             int                 `json:"Amount"`
	Tax                int                 `json:"Tax,omitempty"`
	ClientField1       string              `json:"ClientField1,omitempty"`
	ClientField2       string              `json:"ClientField2,omitempty"`
	ClientField3       string              `json:"ClientField3,omitempty"`
	Overview           string              `json:"Overview,omitempty"`
	Detail             string              `json:"Detail,omitempty"`
	PayMethods         []LinkPlusPayMethod `json:"PayMethods,omitempty"`
	RetURL             string              `json:"RetURL,omitempty"`
	CompleteURL        string              `json:"CompleteURL,omitempty"`
	CancelURL          string              `json:"CancelURL,omitempty"`
	NotifyMailAddress  string              `json:"NotifyMailAddress,omitempty"`
	RetryMax           int                 `json:"RetryMax,omitempty"`
	ExpireDays         int                 `json:"ExpireDays,omitempty"`
	PaymentExpireDate  string              `json:"PaymentExpireDate,omitempty"`
	ResultSkipFlag     int                 `json:"ResultSkipFlag,omitempty"`
	ConfirmSkipFlag    int                 `json:"ConfirmSkipFlag,omitempty"`
	TranDetailShowFlag int                 `json:"TranDetailShowFlag,omitempty"`
}

type getLinkPlusURLRequestJSONCustomer struct {
	MailAddress  string `json:"MailAddress,omitempty"`
	CustomerName string `json:"CustomerName,omitempty"`
	CustomerKana string `json:"CustomerKana,omitempty"`
	TelNo        string `json:"TelNo,omitempty"`
}

type getLinkPlusURLRequestJSONDisplaySetting struct {
	TemplateID   LinkPlusTemplateID   `json:"TemplateID,omitempty"`
	LogoURL      string               `json:"LogoURL,omitempty"`
	ShopName     string               `json:"ShopName,omitempty"`
	ColorPattern LinkPlusColorPattern `json:"ColorPattern,omitempty"`
	Lang         LinkPlusLang         `json:"Lang,omitempty"`
	ShopDomain   string               `json:"ShopDomain,omitempty"`
}

type getLinkPlusURLRequestJSONCreditCard struct {
	JobCD               JobCD  `json:"JobCD,omitempty"`
	Method              int    `json:"Method,omitempty"`
	PayTimes            int    `json:"PayTimes,omitempty"`
	ItemCode            string `json:"ItemCode,omitempty"`
	TdFlag              int    `json:"TdFlag,omitempty"`
	TdTenantName        string `json:"TdTenantName,omitempty"`
	MemberID            string `json:"MemberID,omitempty"`
	SecCodeRequiredFlag int    `json:"SecCodeRequiredFlag,omitempty"`
	SecCodeHiddenFlag   int    `json:"SecCodeHiddenFlag,omitempty"`
	Tds2Type            int    `json:"Tds2Type,omitempty"`
	RegistMemberID      string `json:"RegistMemberID,omitempty"`
	CardInputHiddenFlag int    `json:"CardInputHiddenFlag,omitempty"`
	CardMaxCnt          int    `json:"CardMaxCnt,omitempty"`
}

type getLinkPlusURLResponseJSON struct {
	OrderID     string                                `json:"OrderID"`
	LinkURL     string                                `json:"LinkUrl"`
	ProcessDate string                                `json:"ProcessDate"`
	WarnList    []*getLinkPlusURLResponseJSONWarnList `json:"WarnList"`
	ErrCode     string                                `json:"errCode"`
	ErrInfo     string                                `json:"errInfo"`
}

type getLinkPlusURLResponseJSONWarnList struct {
	WarnCode string `json:"warnCode"`
	WarnInfo string `json:"warnInfo"`
}

func (cli *Client) GetLinkPlusURL(req *GetLinkPlusURLRequest) (*GetLinkPlusURLResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	jsonReq := &getLinkPlusURLRequestJSON{
		GetURLParam: getLinkPlusURLRequestJSONGetURLParam{
			ShopID:                  cli.ShopID,
			ShopPass:                cli.ShopPass,
			GuideMailSendFlag:       req.GuideMailSendFlag,
			ThanksMailSendFlag:      req.ThanksMailSendFlag,
			BccSendMailAddress:      req.BccSendMailAddress,
			FromSendMailAddress:     req.FromSendMailAddress,
			FromSendMailName:        req.FromSendMailName,
			SendMailAddress:         req.SendMailAddress,
			SendMailName:            req.SendMailName,
			CustomerName:            req.CustomerName,
			TemplateNo:              req.TemplateNo,
			GuideSmsSendFlag:        req.GuideSmsSendFlag,
			SMSTelno:                req.SMSTelno,
			AuthenticationQuestion1: req.AuthenticationQuestion1,
			AuthenticationAnswer1:   req.AuthenticationAnswer1,
			AuthenticationQuestion2: req.AuthenticationQuestion2,
			AuthenticationAnswer2:   req.AuthenticationAnswer2,
			AuthenticationRetryMax:  req.AuthenticationRetryMax,
		},
		ConfigID: req.ConfigID,
		Transaction: getLinkPlusURLRequestJSONTransaction{
			OrderID:            req.OrderID,
			Amount:             req.Amount,
			Tax:                req.Tax,
			ClientField1:       req.ClientField1,
			ClientField2:       req.ClientField2,
			ClientField3:       req.ClientField3,
			Overview:           req.Overview,
			Detail:             req.Detail,
			PayMethods:         req.PayMethods,
			RetURL:             req.RetURL,
			CompleteURL:        req.CompleteURL,
			CancelURL:          req.CancelURL,
			NotifyMailAddress:  req.NotifyMailAddress,
			RetryMax:           req.RetryMax,
			ExpireDays:         req.ExpireDays,
			PaymentExpireDate:  req.PaymentExpireDate,
			ResultSkipFlag:     req.ResultSkipFlag,
			ConfirmSkipFlag:    req.ConfirmSkipFlag,
			TranDetailShowFlag: req.TranDetailShowFlag,
		},
		Customer: getLinkPlusURLRequestJSONCustomer{
			MailAddress:  req.CustomerMailAddress,
			CustomerName: req.CustomerName,
			CustomerKana: req.CustomerKana,
			TelNo:        req.TelNo,
		},
		DisplaySetting: getLinkPlusURLRequestJSONDisplaySetting{
			TemplateID:   req.TemplateID,
			LogoURL:      req.LogoURL,
			ShopName:     req.ShopName,
			ColorPattern: req.ColorPattern,
			Lang:         req.Lang,
			ShopDomain:   req.ShopDomain,
		},
	}

	if req.CreditCard != nil {
		jsonReq.CreditCard = getLinkPlusURLRequestJSONCreditCard{
			JobCD:               req.CreditCard.JobCD,
			Method:              req.CreditCard.Method,
			PayTimes:            req.CreditCard.PayTimes,
			ItemCode:            req.CreditCard.ItemCode,
			TdFlag:              req.CreditCard.TdFlag,
			TdTenantName:        req.CreditCard.TdTenantName,
			MemberID:            req.CreditCard.MemberID,
			SecCodeRequiredFlag: req.CreditCard.SecCodeRequiredFlag,
			SecCodeHiddenFlag:   req.CreditCard.SecCodeHiddenFlag,
			Tds2Type:            req.CreditCard.Tds2Type,
			RegistMemberID:      req.CreditCard.RegistMemberID,
			CardInputHiddenFlag: req.CreditCard.CardInputHiddenFlag,
			CardMaxCnt:          req.CreditCard.CardMaxCnt,
		}
	}

	jsonReqBytes, err := json.Marshal(jsonReq)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", cli.APIHost, linkPlusGetUrlPaymentPath),
		bytes.NewBuffer(jsonReqBytes),
	)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")

	httpRes, err := cli.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	bodyBytes, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	jsonRes := &getLinkPlusURLResponseJSON{}
	if err := json.Unmarshal(bodyBytes, jsonRes); err != nil {
		return nil, err
	}

	if jsonRes.ErrCode != "" {
		return &GetLinkPlusURLResponse{
			OrderID: jsonRes.OrderID,
			ErrCode: jsonRes.ErrCode,
			ErrInfo: jsonRes.ErrInfo,
		}, fmt.Errorf("failed to get link plus url: %s", jsonRes.ErrInfo)
	}

	res := &GetLinkPlusURLResponse{
		OrderID:     jsonRes.OrderID,
		LinkURL:     jsonRes.LinkURL,
		ProcessDate: jsonRes.ProcessDate,
		ErrCode:     jsonRes.ErrCode,
		ErrInfo:     jsonRes.ErrInfo,
	}

	warnList := make([]*GetLinkPlusURLResponseWarnList, len(jsonRes.WarnList))
	for i, warn := range jsonRes.WarnList {
		warnList[i] = &GetLinkPlusURLResponseWarnList{
			WarnCode: warn.WarnCode,
			WarnInfo: warn.WarnInfo,
		}
	}
	res.WarnList = warnList

	return res, nil
}
