package remittance

import "fmt"

type ErrorResponse []*ErrorResponseItem

type ErrorResponseItem struct {
	ErrCode string `json:"ErrCode"`
	ErrInfo string `json:"ErrInfo"`
}

func (errResp *ErrorResponse) Error() string {
	msg := "Error:"
	for _, item := range *errResp {
		msg = fmt.Sprintf("%s [%s:%s]", msg, item.ErrCode, item.ErrInfo)
	}
	return msg
}
