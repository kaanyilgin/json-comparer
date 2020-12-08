package model

// DataSet is using for storing data
type DataSet struct {
	objects []JSONObject
}

// NewDataSet creates a new DataSet object
func NewDataSet(data string) *DataSet {
	return &DataSet{
		objects: ParseJSON(data),
	}
}

// Compare compares given dataset
func (dataSet DataSet) Compare(comparedData *DataSet) bool {
	isSameSize := dataSet.getObjectCount() == comparedData.getObjectCount()

	if isSameSize == false {
		return false
	}

	differentObjects := make([]interface{}, 0)

	for objectIndex, object := range dataSet.objects {
		for attributeName, attributeValue := range object.attributes {
			if attributeValue != comparedData.objects[objectIndex].attributes[attributeName] {
				differentObjects = append(differentObjects, attributeValue)
			}
		}
	}

	sameObjectCount := 0

	for mapKey, mapValue := range dataSet.objects {
		for attributeKey := range mapValue.attributes {
			valueFound := false
			for mapKey2, mapValue2 := range comparedData.objects {
				if valueFound {
					break
				}
				for attributeKey2 := range mapValue2.attributes {
					if attributeKey == attributeKey2 {
						attributeValue1 := dataSet.objects[mapKey].attributes[attributeKey]
						attributeValue2 := comparedData.objects[mapKey2].attributes[attributeKey2]
						if attributeValue1 == attributeValue2 {
							sameObjectCount++
							valueFound = true
						}
					}
				}
			}
		}
	}

	if dataSet.getTotalAttributeCount() == comparedData.getTotalAttributeCount() && dataSet.getTotalAttributeCount() == sameObjectCount {
		return true
	}

	return false
}

func (dataSet DataSet) getObjectCount() int {
	return len(dataSet.objects)
}

func (dataSet DataSet) getTotalAttributeCount() int {
	attributeCount := 0

	for i := 0; i < dataSet.getObjectCount(); i++ {
		object := dataSet.objects[i]

		for j := 0; j < len(object.attributes); j++ {
			attributeCount++
		}
	}

	return attributeCount
}
