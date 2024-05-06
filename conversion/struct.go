package conversion

import (
	"bytes"
	"encoding/json"
)

// 將struct轉換為定義的map
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

// 將來源的struct對應至新的struct
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

func StringToMapStringInterface(d string) map[string]interface{} {
	jsonDecoder := json.NewDecoder(bytes.NewReader([]byte(d)))
	jsonDecoder.UseNumber()

	var decodedData map[string]interface{}
	jsonDecoder.Decode(&decodedData)
	return decodedData
}

func MapStringInterfaceToString(d map[string]interface{}) string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
