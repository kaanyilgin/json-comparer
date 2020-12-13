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

// Compare compares the given json object
func (jsonObject JSONObject) GetLength(object IJSONObject) int {
	return len(jsonObject.attributes)
}
