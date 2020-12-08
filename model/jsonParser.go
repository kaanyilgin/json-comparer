package model

import (
	"encoding/json"
)

// ParseJSON parse json into a dynamic object
func ParseJSON(data string) []map[string]interface{} {
	var result []map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	return result
}
