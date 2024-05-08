package main

import (
	"fmt"

	"github.com/abyssparanoia/go-gmo/remittance"
)

func main() {
	cli := newClient()

	res, err := cli.MailDepositRegistration(&remittance.MailDepositRegistrationRequest{
		Method:                 remittance.MailDepositMethodRegister,
		DepositID:              "testdeposit1",
		Amount:                 1000000,
		MailAddress:            "test@dreizehn.com",
		MailDepositAccountName: "ﾀﾅｶﾀﾛｳ",
		Expire:                 "1",
		RemitMethodBank:        remittance.SelectablePaymentMethodEnable,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}
