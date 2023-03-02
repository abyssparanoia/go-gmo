package main

import (
	"fmt"

	"github.com/abyssparanoia/go-gmo/payment"
)

func main() {
	cli := newClient()
	orderID := "orderID"

	// req1 := &payment.EntryTranGANBRequest{
	// 	OrderID: orderID,
	// 	Amount:  2000,
	// }
	// result1, err := cli.EntryTranGANB(req1)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	panic(err)
	// }
	// fmt.Printf("%+v", result1)

	// req2 := &payment.ExecTranGANBRequest{
	// 	OrderID:    orderID,
	// 	AccessID:   result1.AccessID,
	// 	AccessPass: result1.AccessPass,
	// }
	// result2, err := cli.ExecTranGANB(req2)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	panic(err)
	// }
	// fmt.Printf("%+v", result2)

	// req3 := &payment.SearchTradeMultiRequest{
	// 	OrderID: orderID,
	// 	PayType: "36",
	// }
	// result3, err := cli.SearchTradeMulti(req3)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	panic(err)
	// }
	// fmt.Printf("%+v", result3)
	// fmt.Println(result3.GanbBankName)

	// sameMemberReq := &payment.SaveMemberRequest{
	// 	MemberID:   "aadwaddwadwadwada2",
	// 	MemberName: "田中太郎",
	// }

	// saveMemberRes, err := cli.SaveMember(sameMemberReq)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	panic(err)
	// }

	// fmt.Printf("%+v", saveMemberRes)

	// deleteMemberReq := &payment.DeleteMemberRequest{
	// 	MemberID: sameMemberReq.MemberID,
	// }

	// deleteMemberRes, err := cli.DeleteMember(deleteMemberReq)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	panic(err)
	// }

	// fmt.Printf("%+v", deleteMemberRes)

	// sameMemberReq2 := &payment.SaveMemberRequest{
	// 	MemberID:   sameMemberReq.MemberID,
	// 	MemberName: "田中太郎",
	// }

	// saveMemberRes2, err := cli.SaveMember(sameMemberReq2)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	panic(err)
	// }

	// fmt.Printf("%+v", saveMemberRes2)

	paypayEntryTranRes, err := cli.PayPayEntryTran(&payment.PayPayEntryTranRequest{
		OrderID: orderID,
		Amount:  1,
		JobCD:   payment.JobCDAuth,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", paypayEntryTranRes)
}
