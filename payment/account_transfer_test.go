package payment

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	outputFile, err := CreateAccountTranserCSV("test.csv", data)
	defer func() {
		outputFile.Close()
		os.Remove("test.csv")
	}()
	assert.Equal(t, err, nil)

}

func TestCreateAccountTranserFileForSFTP(t *testing.T) {
	t.Skip()
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

	outputFile, err := CreateAccountTranserFileForSFTP("shopID", 1, time.Now(), data)
	defer outputFile.Close()
	assert.Equal(t, err, nil)
}

// func TestCompress(t *testing.T) {
// 	file, err := os.Create("output.tar.gz")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer file.Close()
// 	// set up the gzip writer
// 	gw := gzip.NewWriter(file)
// 	defer gw.Close()
// 	tw := tar.NewWriter(gw)
// 	defer tw.Close()

// 	txtFile, err := os.Open("kfurishopID2004.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer txtFile.Close()

// 	if stat, err := txtFile.Stat(); err == nil {
// 		td := &tar.Header{
// 			Name:    "kfurishopID2004.txt",
// 			Size:    stat.Size(),
// 			Mode:    int64(stat.Mode()),
// 			ModTime: stat.ModTime(),
// 		}
// 		if err := tw.WriteHeader(td); err != nil {
// 			panic(err)
// 		}
// 		// copy the file data to the tarball
// 		if _, err := io.Copy(tw, txtFile); err != nil {
// 			panic(err)
// 		}
// 	}

// }
