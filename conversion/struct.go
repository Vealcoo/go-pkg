package conversion

import (
	"encoding/json"
)

// 將 struct 轉換為 定義的 map
func Struct2Map(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})

	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &data)
	return
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
