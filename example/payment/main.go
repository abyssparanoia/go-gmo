package main

import (
	"fmt"

	"github.com/abyssparanoia/go-gmo/pg/payment"
)

func main() {
	cli := newClient()
	req1 := &payment.EntryTranGANBRequest{
		OrderID: "orderID3",
		Amount:  2000,
	}
	result1, err := cli.EntryTranGANB(req1)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Printf("%+v", result1)

	req2 := &payment.SearchTradeMultiRequest{
		OrderID: "orderID3",
		PayType: "36",
	}
	result2, err := cli.SearchTradeMulti(req2)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Printf("%+v", result2)
}
