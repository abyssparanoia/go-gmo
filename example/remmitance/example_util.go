package main

import (
	"fmt"

	"github.com/abyssparanoia/go-gmo/remittance"
	"github.com/caarlos0/env/v6"
)

type environment struct {
	ShopID   string `env:"SHOP_ID,required"`
	ShopPass string `env:"SHOP_PASSWORD,required"`
}

func newClient() *remittance.Client {
	e := &environment{}
	if err := env.Parse(e); err != nil {
		panic(err)
	}
	cli, err := remittance.NewClient(e.ShopID, e.ShopPass, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("initilize client")
	return cli
}
