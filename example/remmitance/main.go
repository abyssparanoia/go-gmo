package main

import (
	"fmt"

	"github.com/abyssparanoia/go-gmo/remittance"
)

func main() {
	cli := newClient()

	res, err := cli.MailDepositRegistration(&remittance.MailDepositRegistrationRequest{
		Method:                 "1",
		DepositID:              "testdeposit1",
		Amount:                 1000000,
		MailAddress:            "y.sugimoto.paranoia@gmail.com",
		MailDepositAccountName: "ｽｷﾞﾓﾄﾕｳｽｹ",
		Expire:                 "1",
		RemitMethodBank:        "1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}
