package converter

import (
	"bytes"
	"encoding/json"
)

func StructToJsonTagMap(data interface{}) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	// JSONの整形
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(out.String()), &mapData); err != nil {
		return nil, err
	}
	return mapData, err
}
