package main

import (
	"fmt"

	"github.com/abyssparanoia/go-gmo/payment"
	"github.com/caarlos0/env/v6"
)

type environment struct {
	SiteID   string `env:"SITE_ID,required"`
	SitePass string `env:"SITE_PASSWORD,required"`
	ShopID   string `env:"SHOP_ID,required"`
	ShopPass string `env:"SHOP_PASSWORD,required"`
}

func newClient() *payment.Client {
	e := &environment{}
	if err := env.Parse(e); err != nil {
		panic(err)
	}
	cli, err := payment.NewClient(e.SiteID, e.SitePass, e.ShopID, e.ShopPass, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("initilize client")
	return cli
}
