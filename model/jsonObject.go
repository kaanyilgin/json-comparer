package model

// JSONObject stands for json object
type JSONObject struct {
	attributes map[string]interface{}
}

// Compare compares the given json object
func (jsonObject JSONObject) Compare(comparedJSONObject JSONObject) bool {
	return jsonObject.isThereAnyUnMatchedAttribute(comparedJSONObject)
}

func (jsonObject JSONObject) isThereAnyUnMatchedAttribute(comparedJSONObject JSONObject) bool {
	for key, attribute := range jsonObject.attributes {
		for comparedKey, comparedAttribute := range comparedJSONObject.attributes {
			if key == comparedKey && attribute != comparedAttribute {
				return false
			}
		}
	}

	return true
}
