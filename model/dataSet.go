package model

import (
	"strings"
)

// DataSet is using for storing data
type DataSet struct {
	objects []string
}

// NewDataSet creates a new DataSet object
func NewDataSet(data string) *DataSet {
	return &DataSet{
		objects: ParseJSON(data),
	}
}

// Compare compares given dataset
func (dataSet DataSet) Compare(comparedData *DataSet) bool {
	isSameSize := len(dataSet.objects) == len(comparedData.objects)

	if isSameSize == false {
		return false
	}

	differentObjects := make([]string, 0)

	for i := 0; i < len(dataSet.objects); i++ {
		if dataSet.objects[i] != comparedData.objects[i] {
			differentObjects = append(differentObjects, dataSet.objects[i])
		}
	}

	sameObjectCountr := 0

	for i := 0; i < len(dataSet.objects); i++ {
		for j := 0; j < len(comparedData.objects); j++ {
			if dataSet.objects[i] == comparedData.objects[j] {
				sameObjectCountr++
			}
		}
	}

	if len(dataSet.objects) == len(comparedData.objects) && len(dataSet.objects) == sameObjectCountr {
		return true
	}

	return false
}

func parseJSON(data string) []string {
	data = strings.Replace(data, "[", "", 1)
	data = strings.Replace(data, "]", "", 1)
	splittedObject := strings.Split(data, ",")
	return splittedObject
}
