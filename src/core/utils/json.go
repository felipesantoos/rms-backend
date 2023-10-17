package utils

import "encoding/json"

func FormatJSONData(data interface{}) map[string]interface{} {
	var formattedData map[string]interface{}
	dataBytes, _ := json.Marshal(data)
	json.Unmarshal(dataBytes, &formattedData)
	return formattedData
}
