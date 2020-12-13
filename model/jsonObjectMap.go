package model

import (
	"encoding/json"

	"kaanyilgin.com/dataComparer/utility"
)

// JSONObjectMap stands TODO
type JSONObjectMap struct {
	dictionary map[string]*CountJsonObjectPair
}

// NewJSONObject creates a new JSONObject object
func InitJSONObjectMap(attributes map[string]*CountJsonObjectPair) *JSONObjectMap {
	return &JSONObjectMap{
		dictionary: attributes,
	}
}

// Compare compares the given json object
func (jsonObject JSONObjectMap) GetLength() int {
	return len(jsonObject.dictionary)
}

// UnmarshalJSON custom json UnmarshalJSON
func (j *JSONObjectMap) UnmarshalJSON(data []byte) error {

	var unMarshalledData []interface{}
	if err := json.Unmarshal(data, &unMarshalledData); err != nil {
		return err
	}

	j.dictionary = make(map[string]*CountJsonObjectPair, 0)

	for _, object := range unMarshalledData {
		jsonObject2 := object.(map[string]interface{})
		hash := utility.AsSha256(jsonObject2)

		if val, isExist := j.dictionary[hash]; isExist {
			val.jsonObjectCount++
			j.dictionary[hash] = val
		} else {
			j.dictionary[hash] = &CountJsonObjectPair{
				jsonObjectCount: 1,
				jsonObject:      InitJSONObject(jsonObject2),
			}
		}

	}

	return nil
}
