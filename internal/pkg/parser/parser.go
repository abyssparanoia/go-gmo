package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const errCode string = "ErrCode"
const errInfo string = "ErrInfo"

// ParseToMultiObject ... parse to multi object
func ParseToMultiObject(dst interface{}) []map[string]interface{} {
	values := reflect.ValueOf(dst)
	fields := reflect.TypeOf(dst)

	var length int

	tempMap := make(map[string][]string)
	var fieldNames []string

	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i).Interface()
		f := fields.Field(i).Name
		if f == errCode || f == errInfo {
			continue
		}
		fieldNames = append(fieldNames, f)
		slice := strings.Split(fmt.Sprintf("%v", v), "|")
		if length == 0 {
			length = len(slice)
		}
		tempMap[f] = slice
	}

	var multiRes []map[string]interface{}

	for i := 0; i < length; i++ {

		singleRes := make(map[string]interface{})

		for g := 0; g < len(fieldNames); g++ {
			singleRes[fieldNames[g]] = tempMap[fieldNames[g]][i]
		}

		multiRes = append(multiRes, singleRes)
	}

	return multiRes
}

// MapToStruct ... map to struct
func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
