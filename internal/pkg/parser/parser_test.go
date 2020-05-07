package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {

	type response struct {
		OtherField string
		ErrCode    string
		ErrInfo    string
	}

	dst := &response{
		OtherField: "other field value",
		ErrCode:    "E01|E01|E01|E01|E01",
		ErrInfo:    "E01010001|E01020001|E01030002|E01040001|E01060001",
	}

	err := ParseError(dst)
	assert.Equal(t, errors.New("E01010001|E01020001|E01030002|E01040001|E01060001"), err)

}
