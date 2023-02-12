package aozorabank

import "fmt"

type ErrorResponse struct {
	ErrCode    string `json:"errorCode"`
	ErrMessage string `json:"errorMessage"`
}

func (errResp *ErrorResponse) Error() string {
	msg := "Error:"
	msg = fmt.Sprintf("%s [%s:%s]", msg, errResp.ErrCode, errResp.ErrMessage)
	return msg
}
