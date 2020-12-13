package model

// JSONObject stands for json object
type JSONObject struct {
	attributes map[string]interface{}
}

// InitJSONObject creates a new JSONObject object
func InitJSONObject(attributes map[string]interface{}) *JSONObject {
	return &JSONObject{
		attributes,
	}
}

type IJSONObject interface {
	GetLength() int
}

type JSONObjectArray struct {
	dictonary []JSONObject
}

// Compare compares the given json object
func (jsonObjectArray JSONObjectArray) GetLength() int {
	return len(jsonObjectArray.dictonary)
}
