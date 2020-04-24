package payment

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/abyssparanoia/go-gmo/internal/validate"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/gocarina/gocsv"
)

// AccountTransferRecord ... account transfer csv record
type AccountTransferRecord struct {
	SiteID       string `csv:"サイトID" validate:"required"`
	ShopID       string `csv:"ショップID" validate:"required"`
	MemberID     string `csv:"会員ID" validate:"required"`
	Amount       int    `csv:"利用金額" validate:"required"`
	OtherAmount  int    `csv:"税送料"`
	PassbookText string `csv:"通帳記載内容" validate:"required"`
	FreeText     string `csv:"自由項目"`
}

// Validate ... validate record
func (atr *AccountTransferRecord) Validate() error {
	return validate.Struct(atr)
}

// CreateAccountTranserCSV ... create csv file for account transfer
func CreateAccountTranserCSV(
	fileName string,
	data []*AccountTransferRecord,
) (*os.File, error) {

	for _, record := range data {
		if err := record.Validate(); err != nil {
			return nil, err
		}
	}

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(transform.NewWriter(out, japanese.ShiftJIS.NewEncoder()))
		writer.UseCRLF = true
		return gocsv.NewSafeCSVWriter(writer)
	})

	file, err := os.OpenFile(fmt.Sprintf("%s.csv", fileName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err = gocsv.MarshalFile(data, file); err != nil {
		return nil, err
	}

	return file, nil
}
