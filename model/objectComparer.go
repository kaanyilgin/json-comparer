package model

// ObjectComparer provides an interface to compare two object
type ObjectComparer interface {
	Compare(firstJSONObject JSONObject, secondJSONObject JSONObject) (bool, error)
}

//FindAttributeByKeyObjectComparer compares two objects going through whole attribute
type FindAttributeByKeyObjectComparer struct {
	ObjectComparer
}

//Compare compares two object
func (f FindAttributeByKeyObjectComparer) Compare(firstJSONObject JSONObject, secondJSONObject JSONObject) (bool, error) {
	if len(firstJSONObject.attributes) != len(secondJSONObject.attributes) {
		return false, nil
	}

	for key, attribute := range firstJSONObject.attributes {
		comparedValue := secondJSONObject.attributes[key]

		if attribute != comparedValue {
			return false, nil
		}
	}

	return true, nil
}
