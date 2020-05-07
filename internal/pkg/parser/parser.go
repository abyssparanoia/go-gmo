package parser

import (
	"encoding/json"
	"errors"
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
		if f == errCode || f == errInfo || v == "" {
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

// ParseError ... parse error
func ParseError(dst interface{}) error {
	dstMap := StructToMap(dst)
	errInfoField, ok := dstMap[errInfo]
	if !ok {
		panic("no error info field")
	}
	errInfoFieldStr, ok := errInfoField.(string)
	if !ok {
		panic("invalid type")
	}
	if errInfoField == "" {
		return nil
	}
	return errors.New(errInfoFieldStr)
}

// MapToStruct ... convert map to struct
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

// StructToMap ... convert struct to map
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}
