package shiftjis_transformer

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/tomtwinkle/garbledreplacer"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var encoder = japanese.ShiftJIS.NewEncoder()
var decoder = japanese.ShiftJIS.NewDecoder()

var charMap = map[rune]rune{
	'−': '-',  // Convert full-width minus to half-width hyphen
	'〜': '~',  // Convert full-width tilde to half-width tilde
	'－': '-',  // Convert full-width dash to half-width hyphen
	'　': ' ',  // Convert full-width space to half-width space
	'”': '"',  // Convert full-width double quote to half-width
	'“': '"',  // Convert full-width double quote to half-width
	'‘': '\'', // Convert full-width single quote to half-width
	'’': '\'', // Convert full-width single quote to half-width
	'￥': '\\', // Convert full-width yen mark to half-width backslash
	'（': '(',  // Convert full-width left parenthesis to half-width
	'）': ')',  // Convert full-width right parenthesis to half-width
	'［': '[',  // Convert full-width left square bracket to half-width
	'］': ']',  // Convert full-width right square bracket to half-width
	'｛': '{',  // Convert full-width left curly brace to half-width
	'｝': '}',  // Convert full-width right curly brace to half-width
	'＜': '<',  // Convert full-width less than to half-width
	'＞': '>',  // Convert full-width greater than to half-width
	'＝': '=',  // Convert full-width equals to half-width
	'＋': '+',  // Convert full-width plus to half-width
	'；': ';',  // Convert full-width semicolon to half-width
	'：': ':',  // Convert full-width colon to half-width
	'＊': '*',  // Convert full-width asterisk to half-width
	'＆': '&',  // Convert full-width ampersand to half-width
	'％': '%',  // Convert full-width percent to half-width
	'＃': '#',  // Convert full-width hash to half-width
	'＠': '@',  // Convert full-width at sign to half-width
	'！': '!',  // Convert full-width exclamation mark to half-width
	'？': '?',  // Convert full-width question mark to half-width
	'｜': '|',  // Convert full-width vertical bar to half-width
}

// replace specific characters in the string with the mapping table
func replaceUnsupportedChars(input string) string {
	return strings.Map(func(r rune) rune {
		if newR, ok := charMap[r]; ok {
			return newR
		}
		return r
	}, input)
}

func EncodeToShiftJISFromUTF8(input interface{}) error {
	val := reflect.ValueOf(input).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.String {
			original := val.Field(i).String()
			replaced := replaceUnsupportedChars(original)

			var buf bytes.Buffer
			w := transform.NewWriter(&buf, garbledreplacer.NewTransformer(japanese.ShiftJIS, '?'))
			_, err := w.Write([]byte(replaced))
			if err != nil {
				return err
			}
			val.Field(i).SetString(buf.String())
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
