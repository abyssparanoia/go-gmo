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
	TDFlag   TDFlag `schema:"TdFlag,omitempty"`
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

// ExecTranRequest ... exec tran request with member id
type ExecTranRequest = ExecTranRequestWithMemberID

// ExecTranRequestWithMemberID ... exec tran request with member ID
type ExecTranRequestWithMemberID struct {
	AccessID     string `schema:"AccessID" validate:"required"`
	AccessPass   string `schema:"AccessPass" validate:"required"`
	OrderID      string `schema:"OrderID" validate:"required,lte=27"`
	Method       string `schema:"Method,omitempty"`
	PayTimes     int    `schema:"PayTimes"`
	MemberID     string `schema:"MemberID" validate:"required"`
	SeqMode      string `schema:"SeqMode" validate:"required,len=1"`
	CardSeq      int    `schema:"CardSeq" validate:"lte=9999"`
	CardPass     string `schema:"CardPass"`
	SecurityCode string `schema:"SecurityCode"`
	ClientField1 string `schema:"ClientField1,omitempty"`
	ClientField2 string `schema:"ClientField2,omitempty"`
	ClientField3 string `schema:"ClientField3,omitempty"`
}

// Validate ... validate
func (r *ExecTranRequest) Validate() error {
	return validate.Struct(r)
}

// ExecTranRequestWithMemberID ... exec tran request with token
type ExecTranRequestWithToken struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
	OrderID    string `schema:"OrderID" validate:"required,lte=27"`
	Method     string `schema:"Method,omitempty"`
	TokenType  string `schema:"TokenType,omitempty"`
	Token      string `schema:"Token" validate:"required"`
}

// Validate ... validate
func (r *ExecTranRequestWithToken) Validate() error {
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

type ExecTranWith3DSecureRequest struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
	OrderID    string `schema:"OrderID" validate:"required,lte=27"`
	TokenType  string `schema:"TokenType,omitempty"`
	Token      string `schema:"Token" validate:"required"`

	// parameters for 3D Secure
	RetURL               string                  `schema:"RetUrl,omitempty" validate:"required"`
	CallbackType         SecureTran2CallbackType `schema:"CallbackType,omitempty"`
	TDS2ChallengeIndType TDS2ChallengeIndType    `schema:"Tds2ChallengeIndType,omitempty"`
}

func (r *ExecTranWith3DSecureRequest) Validate() error {
	return validate.Struct(r)
}

type ExecTranWith3DSecureResponse struct {
	ACS         string `schema:"ACS"`
	RedirectURL string `schema:"RedirectURL"`
	ErrCode     string `schema:"ErrCode,omitempty"`
	ErrInfo     string `schema:"ErrInfo,omitempty"`
}

// ExecTran ... exec tran with member id
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

// ExecTranWithToken exec tran with token
func (cli *Client) ExecTranWithToken(
	req *ExecTranRequestWithToken,
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

func (cli *Client) ExecTranWith3DSecure(
	req *ExecTranWith3DSecureRequest,
) (*ExecTranWith3DSecureResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &ExecTranWith3DSecureResponse{}
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

type TDS2ResultRequest struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
}

type TDS2ResultResponse struct {
	Tds2TransResult       TDS2TransResult       `schema:"Tds2TransResult"`
	Tds2TransResultReason TDS2TransResultReason `schema:"Tds2TransResultReason"`
	ErrCode               string                `schema:"ErrCode,omitempty"`
	ErrInfo               string                `schema:"ErrInfo,omitempty"`
}

func (r *TDS2ResultRequest) Validate() error {
	return validate.Struct(r)
}

func (cli *Client) TDS2Result(
	req *TDS2ResultRequest,
) (*TDS2ResultResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &TDS2ResultResponse{}
	_, err := cli.do(tds2ResultPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SecureTran2Request struct {
	AccessID   string `schema:"AccessID" validate:"required"`
	AccessPass string `schema:"AccessPass" validate:"required"`
}

// Validate ... validate
func (r *SecureTran2Request) Validate() error {
	return validate.Struct(r)
}

type SecureTran2Response struct {
	OrderID      string `schema:"OrderID,omitempty"`
	Forward      string `schema:"Forward,omitempty"`
	Method       string `schema:"Method,omitempty"`
	PayTimes     string `schema:"PayTimes,omitempty"`
	Approve      string `schema:"Approve,omitempty"`
	TranID       string `schema:"TranID,omitempty"`
	TranDate     string `schema:"TranDate,omitempty"`
	CheckString  string `schema:"CheckString,omitempty"`
	ClientField1 string `schema:"ClientField1,omitempty"`
	ClientField2 string `schema:"ClientField2,omitempty"`
	ClientField3 string `schema:"ClientField3,omitempty"`
}

func (cli *Client) SecureTran2(
	req *SecureTran2Request,
) (*SecureTran2Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &SecureTran2Response{}
	_, err := cli.do(secureTran2Path, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
