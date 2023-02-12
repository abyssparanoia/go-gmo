package aozorabank

import "fmt"

type ErrorResponse struct {
	ErrCode    string `json:"errorCode"`
	ErrMessage string `json:"errorMessage"`
}

func (errResp *ErrorResponse) Error() string {
	return fmt.Sprintf("Error: [%s:%s]", errResp.ErrCode, errResp.ErrMessage)
}
