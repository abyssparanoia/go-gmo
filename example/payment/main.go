package main

import (
	"fmt"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/abyssparanoia/go-gmo/pg/payment"
)

func main() {
	cli := newClient()
	orderID := "orderID5"

	req1 := &payment.EntryTranGANBRequest{
		OrderID: orderID,
		Amount:  2000,
	}
	result1, err := cli.EntryTranGANB(req1)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Printf("%+v", result1)

	req2 := &payment.ExecTranGANBRequest{
		OrderID:    orderID,
		AccessID:   result1.AccessID,
		AccessPass: result1.AccessPass,
	}
	result2, err := cli.ExecTranGANB(req2)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Printf("%+v", result2)

	req3 := &payment.SearchTradeMultiRequest{
		OrderID: orderID,
		PayType: "36",
	}
	result3, err := cli.SearchTradeMulti(req3)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Printf("%+v", result3)
	fmt.Println(result3.GanbBankName)

	str, _, err := transform.String(japanese.ShiftJIS.NewDecoder(), result3.GanbBankName)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Println(str)
}
