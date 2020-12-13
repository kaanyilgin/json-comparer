package model

import (
	"encoding/json"

	"kaanyilgin.com/dataComparer/utility"
)

// JSONObjectMap stands TODO
type JSONObjectMap struct {
	dictionary map[string]*JSONObject
}

// NewJSONObject creates a new JSONObject object
func InitJSONObjectMap(attributes map[string]*JSONObject) *JSONObjectMap {
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

	j.dictionary = make(map[string]*JSONObject, 0)

	for _, object := range unMarshalledData {
		jsonObject2 := object.(map[string]interface{})
		hash := utility.AsSha256(jsonObject2)
		j.dictionary[hash] = InitJSONObject(jsonObject2)
	}

	return nil
}
