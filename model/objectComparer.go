package model

// ObjectComparer provides an interface to compare two object
type ObjectComparer interface {
	Compare(firstJSONObject JSONObject, secondJSONObject JSONObject) bool
}

//FindAttributeByKeyObjectComparer compares two objects going through whole attribute
type FindAttributeByKeyObjectComparer struct {
}

//Compare compares two object
func (f FindAttributeByKeyObjectComparer) Compare(firstJSONObject JSONObject, secondJSONObject JSONObject) bool {
	if len(firstJSONObject.attributes) != len(secondJSONObject.attributes) {
		return false
	}

	for key, attribute := range firstJSONObject.attributes {
		comparedValue := secondJSONObject.attributes[key]

		if attribute != comparedValue {
			return false
		}
	}

	return true
}
