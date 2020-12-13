package model

import (
	"encoding/json"

	"kaanyilgin.com/dataComparer/utility"
)

// JSONObjectMap stores json object as hash map and count json object pair type
type JSONObjectMap struct {
	dictionary map[string]*CountJSONObjectPair
}

// InitJSONObjectMap is a constructor for JSONObjectMap
func InitJSONObjectMap(attributes map[string]*CountJSONObjectPair) *JSONObjectMap {
	return &JSONObjectMap{
		dictionary: attributes,
	}
}

// GetLength returns the map length
func (j JSONObjectMap) GetLength() int {
	return len(j.dictionary)
}

// UnmarshalJSON parse json to JSONObjectMap
func (j *JSONObjectMap) UnmarshalJSON(data []byte) error {

	var unMarshalledData []interface{}
	if err := json.Unmarshal(data, &unMarshalledData); err != nil {
		return err
	}

	j.dictionary = make(map[string]*CountJSONObjectPair, 0)

	for _, object := range unMarshalledData {
		jsonObject2 := object.(map[string]interface{})
		hash := utility.AsSha256(jsonObject2)

		if val, isExist := j.dictionary[hash]; isExist {
			val.jsonObjectCount++
			j.dictionary[hash] = val
		} else {
			j.dictionary[hash] = &CountJSONObjectPair{
				jsonObjectCount: 1,
				jsonObject:      InitJSONObject(jsonObject2),
			}
		}

	}

	return nil
}
