package model

import (
	"encoding/json"
)

// JSONParser provides an interface to parse json to JSONObjects
type JSONParser interface {
	ParseJSON(data string) JSONObjectMap
}

//SystemJSONParser parse json to JSONObjects by using "encoding/json"
type SystemJSONParser struct {
}

func newSystemJSONParser() *SystemJSONParser {
	return &SystemJSONParser{}
}

// ParseJSON parse json into a dynamic object
func (s SystemJSONParser) ParseJSON(data string) JSONObjectMap {
	var result JSONObjectMap
	json.Unmarshal([]byte(data), &result)
	return result
}
