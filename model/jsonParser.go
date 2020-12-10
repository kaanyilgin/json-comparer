package model

import (
	"encoding/json"
)

// ParseJSON parse json into a dynamic object
func ParseJSON(data string, objectComparer ObjectComparer) []JSONObject {
	var result []map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	objects := make([]JSONObject, 0)

	for _, attributes := range result {
		object := NewJSONObject(attributes, objectComparer)
		objects = append(objects, *object)
	}

	return objects
}
