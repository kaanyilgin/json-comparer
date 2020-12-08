package model

// JSONObject stands for json object
type JSONObject struct {
	attributes map[string]interface{}
}

// Compare compares the given json object
func (jsonObject JSONObject) Compare(comparedJSONObject JSONObject) bool {
	return false
}
