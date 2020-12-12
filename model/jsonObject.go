package model

// JSONObject stands for json object
type JSONObject struct {
	attributes map[string]interface{}
}

// NewJSONObject creates a new JSONObject object
func NewJSONObject(attributes map[string]interface{}) *JSONObject {
	return &JSONObject{
		attributes,
	}
}

// Compare compares the given json object
func (jsonObject JSONObject) Compare(objectComparer ObjectComparer, comparedJSONObject JSONObject) (bool, error) {
	return objectComparer.Compare(jsonObject, comparedJSONObject)
}
