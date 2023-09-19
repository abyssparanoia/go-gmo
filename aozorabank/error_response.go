package aozorabank

import (
	"fmt"
)

type ErrorResponse struct {
	ErrorCode            string                `json:"errorCode"`
	ErrorMessage         string                `json:"errorMessage"`
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
	return fmt.Sprintf("Error: [%s:%s], ErrorDetails: %v, TransferErrorDetails: %v", errResp.ErrorCode, errResp.ErrorMessage, errResp.ErrorDetails, errResp.TransferErrorDetails)
}

type AuthErrorResponse struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
}

func (errResp *AuthErrorResponse) Error() string {
	return fmt.Sprintf("Error: [%s:%s],  ErrorURI: %s", errResp.ErrorCode, errResp.ErrorDescription, errResp.ErrorURI)
}

type ErrorCode string

const (
	ErrorCodeUnderMaintenance ErrorCode = "WG_ERR_300"
)

func (e ErrorCode) String() string {
	return string(e)
}
