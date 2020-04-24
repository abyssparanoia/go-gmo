package payment

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateAccountTranserCSV(t *testing.T) {

	data := []*AccountTransferRecord{
		{
			SiteID:       "siteID",
			ShopID:       "shopID",
			MemberID:     "memberID",
			Amount:       10000,
			OtherAmount:  1000,
			PassbookText: "ｱｲｳｴｵ",
			FreeText:     "free text",
		},
		{SiteID: "siteID",
			ShopID:       "shopID",
			MemberID:     "memberID",
			Amount:       10000,
			OtherAmount:  1000,
			PassbookText: "ｱｲｳｴｵ",
			FreeText:     "free text",
		},
	}

	outputFile, err := CreateAccountTranserCSV("test", data)
	defer outputFile.Close()
	assert.Equal(t, err, nil)

}
