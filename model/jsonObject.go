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

// IJSONObject is an interface for different data set structs
type IJSONObject interface {
	GetLength() int
}

// JSONObjectArray stores the json objects in a plain array
type JSONObjectArray struct {
	dictonary []JSONObject
}

// GetLength returns the array length
func (jsonObjectArray JSONObjectArray) GetLength() int {
	return len(jsonObjectArray.dictonary)
}
