package remittance

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/converter"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

type MailDepositRegistrationRequest struct {
	Method                 Method                  `json:"Method" validate:"required"`
	DepositID              string                  `json:"Deposit_ID" validate:"required,max=27"`
	Amount                 int                     `json:"Amount" validate:"required"`
	MailAddress            string                  `json:"Mail_Address" validate:"required,email,max=200"`
	ShopMailAddress        string                  `json:"Shop_Mail_Address,omitempty"`
	MailDepositAccountName string                  `json:"Mail_Deposit_Account_Name,omitempty"`
	Expire                 string                  `json:"Expire,omitempty"`
	AuthCode               string                  `json:"Auth_Code,omitempty"`
	AuthCode2              string                  `json:"Auth_Code2,omitempty"`
	AuthCode3              string                  `json:"Auth_Code3,omitempty"`
	RemitMethodBank        SelectablePaymentMethod `json:"Remit_Method_Bank,omitempty"`
	RemitMethodSevenatm    SelectablePaymentMethod `json:"Remit_Method_Sevenatm,omitempty"`
	SevenatmPaymentTermDay int                     `json:"Sevenatm_Payment_Term_Day,omitempty"`
	RemitMethodAmazongift  SelectablePaymentMethod `json:"Remit_Method_Amazongift,omitempty"`
	RemitMethodAupay       SelectablePaymentMethod `json:"Remit_Method_Aupay,omitempty"`
	MailTemplateFree1      string                  `json:"Mail_Template_Free1,omitempty"`
	MailTemplateFree2      string                  `json:"Mail_Template_Free2,omitempty"`
	MailTemplateFree3      string                  `json:"Mail_Template_Free3,omitempty"`
	MailTemplateNumber     int                     `json:"Mail_Template_Number,omitempty"`
	BankID                 string                  `json:"Bank_ID,omitempty"`
}

func (r *MailDepositRegistrationRequest) Validate() error {
	return validate.Struct(r)
}

type MailDepositRegistrationResponse struct {
	DepositID string `json:"Deposit_ID"`
	Method    string `json:"Method"`
	Amount    string `json:"Amount"`
	Expire    string `json:"Expire"`
}

func (cli *Client) MailDepositRegistration(
	req *MailDepositRegistrationRequest,
) (*MailDepositRegistrationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqMap, err := converter.StructToJsonTagMap(req)
	if err != nil {
		return nil, err
	}
	res := &MailDepositRegistrationResponse{}
	_, err = cli.do(mailDepositRegistrationPath, reqMap, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
