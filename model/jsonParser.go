package model

import (
	"encoding/json"
)

// JSONParser provides an interface to parse json to JSONObjects
type JSONParser interface {
	ParseJSON(data string) interface{}
}

// JSONObjectMapParser parse json to JSONObjects by using "encoding/json"
type JSONObjectMapParser struct {
}

// ParseJSON parse json into a dynamic object
func (j JSONObjectMapParser) ParseJSON(data string) interface{} {
	var result JSONObjectMap
	json.Unmarshal([]byte(data), &result)
	return result
}

// JSONObjectParser parse json to JSONObjects by using "encoding/json"
type JSONObjectParser struct {
}

// ParseJSON parse json into a dynamic object
func (j JSONObjectParser) ParseJSON(data string) interface{} {
	var result []map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	return result
}
