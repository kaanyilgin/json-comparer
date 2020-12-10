package model

// ObjectComparer provides an interface to compare two object
type ObjectComparer interface {
	Compare(firstJSONObject JSONObject, secondJSONObject JSONObject) (bool, error)
}

//LoopingTwoObjectAttributeComparer compares two objects going through whole attribute
type LoopingTwoObjectAttributeComparer struct {
	DataSetComparer
}

//Compare compares two object
func (l LoopingTwoObjectAttributeComparer) Compare(firstJSONObject JSONObject, secondJSONObject JSONObject) (bool, error) {
	if len(firstJSONObject.attributes) != len(secondJSONObject.attributes) {
		return false, nil
	}

	for key, attribute := range firstJSONObject.attributes {
		for comparedKey, comparedAttribute := range secondJSONObject.attributes {
			if key == comparedKey && attribute != comparedAttribute {
				return false, nil
			}
		}
	}

	return true, nil
}
