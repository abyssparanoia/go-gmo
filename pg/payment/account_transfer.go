package payment

import (
	"archive/tar"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"

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

func createAccountTranserFile(
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

	file, err := os.OpenFile(fmt.Sprintf("%s", fileName), os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		file.Close()
		return nil, err
	}

	if err = gocsv.MarshalFile(data, file); err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

// CreateAccountTranserCSV ... create csv file for account transfer
func CreateAccountTranserCSV(
	fileName string,
	data []*AccountTransferRecord,
) (*os.File, error) {
	return createAccountTranserFile(fileName, data)
}

// CreateAccountTranserFileForSFTP ... create tar file for sftp
func CreateAccountTranserFileForSFTP(
	shopID string,
	sequence int,
	uploadTime time.Time,
	data []*AccountTransferRecord,
) (*os.File, error) {

	year := strconv.Itoa(uploadTime.Year())[2:]

	orgMonth := int(uploadTime.Month())
	var month string
	if orgMonth < 10 {
		month = fmt.Sprintf("0%s", strconv.Itoa(orgMonth))
	} else {
		month = strconv.Itoa(orgMonth)
	}

	var sequenceStr string
	if sequence < 10 {
		sequenceStr = fmt.Sprintf("0%d", sequence)
	} else {
		sequenceStr = fmt.Sprintf("%d", sequence)
	}

	baseFileName := fmt.Sprintf("kfuri%s%s%s%s", shopID, year, month, sequenceStr)
	txtFileName := fmt.Sprintf("%s.txt", baseFileName)
	gzFileName := fmt.Sprintf("%s.tar.gz", baseFileName)

	txtFile, err := createAccountTranserFile(txtFileName, data)
	if err != nil {
		return nil, err
	}
	txtFile.Close()

	txtFile, err = os.Open(txtFileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		txtFile.Close()
		os.Remove(txtFileName)
	}()

	stat, err := txtFile.Stat()
	if err != nil {
		return nil, err
	}

	outputFile, err := os.Create(gzFileName)
	if err != nil {
		return nil, err
	}

	gzw := gzip.NewWriter(outputFile)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	hdr := &tar.Header{
		Name:    txtFileName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}
	if err = tw.WriteHeader(hdr); err != nil {
		outputFile.Close()
		return nil, err
	}

	if _, err = io.Copy(tw, txtFile); err != nil {
		outputFile.Close()
		return nil, err
	}

	return outputFile, nil
}
