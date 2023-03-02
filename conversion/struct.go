package conversion

import (
	"bytes"
	"encoding/json"
)

// 將 struct 轉換為定義的 map
func Struct2Map(data interface{}) map[string]interface{} {
	var dataMap map[string]interface{}

	if marshalData, err := json.Marshal(data); err != nil {
		return nil
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalData))
		d.UseNumber() // float64 to number (避免科學記號)
		if err := d.Decode(&dataMap); err != nil {
			return nil

		} else {
			for k, v := range dataMap {
				dataMap[k] = v
			}
		}
	}

	return dataMap
}

// 將 來源的 struct 對應至新的 struct
func Struct2Struct(source interface{}, target interface{}) error {
	dJson, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dJson, target)
	if err != nil {
		return err
	}

	return nil
}
