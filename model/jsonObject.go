package model

// JSONObject stands for json object
type JSONObject struct {
	attributes     map[string]interface{}
	objectComparer ObjectComparer
}

// NewJSONObject creates a new NewJSONObject object
func NewJSONObject(attributes map[string]interface{}, objectComparer ObjectComparer) *JSONObject {
	return &JSONObject{
		attributes:     attributes,
		objectComparer: objectComparer,
	}
}

// Compare compares the given json object
func (jsonObject JSONObject) Compare(comparedJSONObject JSONObject) (bool, error) {
	return jsonObject.objectComparer.Compare(jsonObject, comparedJSONObject)
}
