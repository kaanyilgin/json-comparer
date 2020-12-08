package model

import (
	"encoding/json"
)

// ParseJSON parse json into a dynamic object
func ParseJSON(data string) []JSONObject {
	var result []map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	objects := make([]JSONObject, 0)

	for mapKey := range result {
		object := &JSONObject{
			attributes: result[mapKey],
		}

		objects = append(objects, *object)
	}

	return objects
}
