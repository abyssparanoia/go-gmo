package shiftjis_transformer

import (
	"encoding/json"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var encoder = japanese.ShiftJIS.NewEncoder()
var decoder = japanese.ShiftJIS.NewDecoder()

func EncodeToShiftJISFromUTF8(input interface{}) error {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	encodedBytes, _, err := transform.Bytes(encoder, inputBytes)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(encodedBytes, input); err != nil {
		return err
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
