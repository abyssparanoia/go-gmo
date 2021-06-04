package shiftjis_transformer

import (
	"reflect"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var encoder = japanese.ShiftJIS.NewEncoder()
var decoder = japanese.ShiftJIS.NewDecoder()

func EncodeToShiftJISFromUTF8(input interface{}) error {
	val := reflect.ValueOf(input).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.String {
			encoded, _, err := transform.String(encoder, val.Field(i).String())
			if err != nil {
				return err
			}
			val.Field(i).SetString(encoded)
		}
	}
	return nil
}

func DecodeToUTF8FromShiftJIS(inputBytes []byte) ([]byte, error) {
	decodedBytes, _, err := transform.Bytes(decoder, inputBytes)
	if err != nil {
		return nil, err
	}
	return decodedBytes, nil
}
