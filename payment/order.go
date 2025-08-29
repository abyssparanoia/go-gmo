// nolint:nilness
package payment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type InquiryOrderRequest struct {
	OrderID string `validate:"required"`
}

func (r *InquiryOrderRequest) Validate() error {
	return validate.Struct(r)
}

type InquiryOrderResponse struct {
	OrderReference       *InquiryOrderResponseOrderReference
	CreditResult         *InquiryOrderResponseCreditResult
	Tds2Result           *InquiryOrderResponseTds2Result
	FraudDetectionResult *InquiryOrderResponseFraudDetectionResult
	WalletResult         *InquiryOrderResponseWalletResult
	CashResult           *InquiryOrderResponseCashResult
}

type InquiryOrderResponseOrderReference struct {
	AccessID   string
	AccessPass string
	OrderID    string
	Status     TradeMultiStatus
	Created    time.Time
	Updated    time.Time
}

type InquiryOrderResponseClientFields struct {
	ClientField1 string
	ClientField2 string
	ClientField3 string
}

type InquiryOrderResponseCardResult struct {
	CardNumber     string
	CardholderName string
	ExpiryMonth    string
	ExpiryYear     string
	IssuerCode     string
}

type InquiryOrderResponseCreditResult struct {
	CardType              string
	CardResult            *InquiryOrderResponseCardResult
	ForwardedAcquirerCode string
	ApprovalCode          string
	NwTransactionID       string
	TransactionDateTime   time.Time
	CaptureExpiryDateTime time.Time
	UseTds2               bool
	UseFraudDetection     bool
}

type InquiryOrderResponseTds2Result struct {
	Eci                   string
	RequiresChallenge     bool
	Tds2TransResult       string
	Tds2TransResultReason string
}

type InquiryOrderResponseFraudDetectionResult struct {
	ScreeningType          string
	ScreeningTransactionID string
	ScreeningResultCode    string
	ScreeningResultRawData string
}

type InquiryOrderResponseWalletResult struct {
	WalletType     string
	SettlementCode string
}

type InquiryOrderResponseCashResult struct {
	CashType                  string
	PaymentExpiryDateTime     time.Time
	PayEasyPaymentInformation *InquiryOrderResponsePayEasyPaymentInformation
}

type InquiryOrderResponsePayEasyPaymentInformation struct {
	CustomerNumber     string
	InstitutionCode    string
	ConfirmationNumber string
	BankURL            string
}

type inquiryOrderRequestJSON struct {
	OrderID string `json:"orderId"`
}

type inquiryOrderResponseJSON struct {
	OrderReference *struct {
		AccessID     string           `json:"accessId"`
		AccessPass   string           `json:"accessPass"`
		OrderID      string           `json:"orderId"`
		Status       TradeMultiStatus `json:"status"`
		Created      time.Time        `json:"created"`
		Updated      time.Time        `json:"updated"`
		Amount       string           `json:"amount"`
		ClientFields struct {
			ClientField1 string `json:"clientField1"`
			ClientField2 string `json:"clientField2"`
			ClientField3 string `json:"clientField3"`
		} `json:"clientFields"`
		ChargeType string `json:"chargeType"`
	} `json:"orderReference,omitempty"`
	CreditResult *struct {
		CardType   string `json:"cardType"`
		CardResult struct {
			CardNumber     string `json:"cardNumber"`
			CardholderName string `json:"cardholderName"`
			ExpiryMonth    string `json:"expiryMonth"`
			ExpiryYear     string `json:"expiryYear"`
			IssuerCode     string `json:"issuerCode"`
		} `json:"cardResult"`
		ForwardedAcquirerCode string    `json:"forwardedAcquirerCode"`
		ApprovalCode          string    `json:"approvalCode"`
		NwTransactionID       string    `json:"nwTransactionId"`
		TransactionDateTime   time.Time `json:"transactionDateTime"`
		CaptureExpiryDateTime time.Time `json:"captureExpiryDateTime"`
		UseTds2               bool      `json:"useTds2"`
		UseFraudDetection     bool      `json:"useFraudDetection"`
	} `json:"creditResult,omitempty"`
	Tds2Result *struct {
		Eci                   string `json:"eci"`
		RequiresChallenge     bool   `json:"requiresChallenge"`
		Tds2TransResult       string `json:"tds2TransResult"`
		Tds2TransResultReason string `json:"tds2TransResultReason"`
	} `json:"tds2Result,omitempty"`
	FraudDetectionResult *struct {
		ScreeningType          string `json:"screeningType"`
		ScreeningTransactionID string `json:"screeningTransactionId"`
		ScreeningResultCode    string `json:"screeningResultCode"`
		ScreeningResultRawData string `json:"screeningResultRawData"`
	} `json:"fraudDetectionResult,omitempty"`
	WalletResult *struct {
		WalletType     string `json:"walletType"`
		SettlementCode string `json:"settlementCode"`
	} `json:"walletResult,omitempty"`
	CashResult *struct {
		CashType                  string    `json:"cashType"`
		PaymentExpiryDateTime     time.Time `json:"paymentExpiryDateTime"`
		PayEasyPaymentInformation struct {
			CustomerNumber     string `json:"customerNumber"`
			InstitutionCode    string `json:"institutionCode"`
			ConfirmationNumber string `json:"confirmationNumber"`
			BankURL            string `json:"bankUrl"`
		} `json:"pay-easyPaymentInformation"`
	} `json:"cashResult,omitempty"`
}

type inquiryOrderErrorResponseJSON struct {
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func (cli *Client) InquiryOrder(req *InquiryOrderRequest) (*InquiryOrderResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	jsonReq := &inquiryOrderRequestJSON{
		OrderID: req.OrderID,
	}

	jsonRes := &inquiryOrderResponseJSON{}
	errResp := &inquiryOrderErrorResponseJSON{}

	httpRes, err := cli.doOpenAPI(inquiryOrderPath, jsonReq, jsonRes, errResp)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to inquiry order: %s", errResp.Detail)
	}

	res := &InquiryOrderResponse{}
	if jsonRes.OrderReference != nil {
		res.OrderReference = &InquiryOrderResponseOrderReference{
			AccessID:   jsonRes.OrderReference.AccessID,
			AccessPass: jsonRes.OrderReference.AccessPass,
			OrderID:    jsonRes.OrderReference.OrderID,
			Status:     jsonRes.OrderReference.Status,
			Created:    jsonRes.OrderReference.Created,
			Updated:    jsonRes.OrderReference.Updated,
		}
	}
	if jsonRes.CreditResult != nil {
		res.CreditResult = &InquiryOrderResponseCreditResult{
			CardType: jsonRes.CreditResult.CardType,
			CardResult: &InquiryOrderResponseCardResult{
				CardNumber:     jsonRes.CreditResult.CardResult.CardNumber,
				CardholderName: jsonRes.CreditResult.CardResult.CardholderName,
				ExpiryMonth:    jsonRes.CreditResult.CardResult.ExpiryMonth,
				ExpiryYear:     jsonRes.CreditResult.CardResult.ExpiryYear,
				IssuerCode:     jsonRes.CreditResult.CardResult.IssuerCode,
			},
			ForwardedAcquirerCode: jsonRes.CreditResult.ForwardedAcquirerCode,
			ApprovalCode:          jsonRes.CreditResult.ApprovalCode,
			NwTransactionID:       jsonRes.CreditResult.NwTransactionID,
			TransactionDateTime:   jsonRes.CreditResult.TransactionDateTime,
			CaptureExpiryDateTime: jsonRes.CreditResult.CaptureExpiryDateTime,
			UseTds2:               jsonRes.CreditResult.UseTds2,
			UseFraudDetection:     jsonRes.CreditResult.UseFraudDetection,
		}
	}
	if jsonRes.Tds2Result != nil {
		res.Tds2Result = &InquiryOrderResponseTds2Result{
			Eci:                   jsonRes.Tds2Result.Eci,
			RequiresChallenge:     jsonRes.Tds2Result.RequiresChallenge,
			Tds2TransResult:       jsonRes.Tds2Result.Tds2TransResult,
			Tds2TransResultReason: jsonRes.Tds2Result.Tds2TransResultReason,
		}
	}
	if jsonRes.FraudDetectionResult != nil {
		res.FraudDetectionResult = &InquiryOrderResponseFraudDetectionResult{
			ScreeningType:          jsonRes.FraudDetectionResult.ScreeningType,
			ScreeningTransactionID: jsonRes.FraudDetectionResult.ScreeningTransactionID,
			ScreeningResultCode:    jsonRes.FraudDetectionResult.ScreeningResultCode,
			ScreeningResultRawData: jsonRes.FraudDetectionResult.ScreeningResultRawData,
		}
	}
	if jsonRes.WalletResult != nil {
		res.WalletResult = &InquiryOrderResponseWalletResult{
			WalletType:     jsonRes.WalletResult.WalletType,
			SettlementCode: jsonRes.WalletResult.SettlementCode,
		}
	}
	if jsonRes.CashResult != nil {
		res.CashResult = &InquiryOrderResponseCashResult{
			CashType:              jsonRes.CashResult.CashType,
			PaymentExpiryDateTime: jsonRes.CashResult.PaymentExpiryDateTime,
			PayEasyPaymentInformation: &InquiryOrderResponsePayEasyPaymentInformation{
				CustomerNumber:     jsonRes.CashResult.PayEasyPaymentInformation.CustomerNumber,
				InstitutionCode:    jsonRes.CashResult.PayEasyPaymentInformation.InstitutionCode,
				ConfirmationNumber: jsonRes.CashResult.PayEasyPaymentInformation.ConfirmationNumber,
				BankURL:            jsonRes.CashResult.PayEasyPaymentInformation.BankURL,
			},
		}
	}

	return res, nil
}
