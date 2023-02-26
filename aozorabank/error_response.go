package aozorabank

import "fmt"

type ErrorResponse struct {
	ErrCode              string                `json:"errorCode"`
	ErrMessage           string                `json:"errorMessage"`
	ErrorDetails         []ErrorDetail         `json:"errorDetails"`
	TransferErrorDetails []TransferErrorDetail `json:"transferErrorDetails"`
}

type ErrorDetail struct {
	ErrorDetailsCode   string `json:"errorDetailsCode"`
	ErrorDetailMessage string `json:"errorDetailsMessage"`
}

type TransferErrorDetail struct {
	ItemID       string        `json:"itemId"`
	ErrorDetails []ErrorDetail `json:"errorDetails"`
}

func (errResp *ErrorResponse) Error() string {
	return fmt.Sprintf("Error: [%s:%s], ErrorDetails: %#v, TransferErrorDetails: %#v", errResp.ErrCode, errResp.ErrMessage, errResp.ErrorDetails, errResp.TransferErrorDetails)
}
