package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// EntryTranRequest ... entry tran request type
type EntryTranRequest struct {
	OrderID  string `schema:"OrderID" validate:"required,lte=27"`
	JobCD    JobCD  `schema:"JobCd" validate:"required"`
	ItemCode string `schema:"ItemCode,omitempty"`
	Amount   int    `schema:"Amount,omitempty"`
	Tax      int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *EntryTranRequest) Validate() error {
	return validate.Struct(r)
}

// EntryTranResponse ... entry tran response
type EntryTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// EntryTran ... entry tran
func (cli *Client) EntryTran(
	req *EntryTranRequest,
) (*EntryTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &EntryTranResponse{}
	_, err := cli.do(entryTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ExecTranRequest ... exec tran request
type ExecTranRequest struct {
	AccessID     string `schema:"AccessID" validate:"required"`
	AccessPass   string `schema:"AccessPass" validate:"required"`
	OrderID      string `schema:"OrderID" validate:"required,lte=27"`
	Method       string `schema:"Method,omitempty"`
	PayTimes     int    `schema:"PayTimes"`
	MemberID     string `schema:"MemberID" validate:"required"`
	SeqMode      string `schema:"SeqMode" validate:"required,len=1"`
	CardSeq      int    `schema:"CardSeq" validate:"lte=4"`
	CardPass     string `schema:"CardPass"`
	SecurityCode string `schema:"SecurityCode"`
}

// Validate ... validate
func (r *ExecTranRequest) Validate() error {
	return validate.Struct(r)
}

// ExecTranResponse ... exec tran response
type ExecTranResponse struct {
	ACS         string `schema:"ACS"`
	OrderID     string `schema:"OrderID"`
	Forward     string `schema:"forward"`
	Method      string `schema:"Method"`
	PayTimes    string `schema:"PayTimes"`
	Approve     string `schema:"Approve"`
	TranID      string `schema:"TranID"`
	TranDate    string `schema:"TranDate"`
	CheckString string `schema:"CheckString"`
	ACSUrl      string `schema:"ACSUrl"`
	PaReq       string `schema:"PaReq"`
	MD          string `schema:"MD"`
	ErrCode     string `schema:"ErrCode,omitempty"`
	ErrInfo     string `schema:"ErrInfo,omitempty"`
}

// ExecTran ... exec tran
func (cli *Client) ExecTran(
	req *ExecTranRequest,
) (*ExecTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &ExecTranResponse{}
	_, err := cli.do(execTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AlterTranRequest ... alter tran request
type AlterTranRequest struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
	JobCD      JobCD  `schema:"JobCd" validate:"required"`
	Amount     int    `schema:"Amount,omitempty"`
	Tax        int    `schema:"Tax,omitempty"`
	Method     string `schema:"Method,omitempty"`
	PayTimes   int    `schema:"PayTimes,omitempty"`
}

// Validate ... validate
func (r *AlterTranRequest) Validate() error {
	return validate.Struct(r)
}

// AlterTranResponse ... alter tran response
type AlterTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	Forward    string `schema:"Forward,omitempty"`
	Approve    string `schema:"Approve,omitempty"`
	TranID     string `schema:"TranID,omitempty"`
	TranDate   string `schema:"TranDate,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// AlterTran ... alter tran
func (cli *Client) AlterTran(
	req *AlterTranRequest,
) (*AlterTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &AlterTranResponse{}
	_, err := cli.do(alterTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChangeTranRequest ... change tran request
type ChangeTranRequest struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
	JobCD      JobCD  `schema:"JobCd" validate:"required"`
	Amount     int    `schema:"Amount" validate:"required"`
	Tax        int    `schema:"Tax,omitempty"`
}

// Validate ... validate
func (r *ChangeTranRequest) Validate() error {
	return validate.Struct(r)
}

// ChangeTranResponse ... change tran response
type ChangeTranResponse struct {
	AccessID   string `schema:"AccessID,omitempty"`
	AccessPass string `schema:"AccessPass,omitempty"`
	Forward    string `schema:"Forward,omitempty"`
	Approve    string `schema:"Approve,omitempty"`
	TranID     string `schema:"TranID,omitempty"`
	TranDate   string `schema:"TranDate,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// ChangeTran ... change tran
func (cli *Client) ChangeTran(
	req *ChangeTranRequest,
) (*ChangeTranResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &ChangeTranResponse{}
	_, err := cli.do(changeTranPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
