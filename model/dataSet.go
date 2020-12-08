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
func (dataSet DataSet) Compare(data string) bool {
	objects2 := parseJSON(data)
	isSameSize := len(dataSet.objects) == len(objects2)

	if isSameSize == false {
		return false
	}

	differentObjects := make([]string, 0)

	for i := 0; i < len(dataSet.objects); i++ {
		if dataSet.objects[i] != objects2[i] {
			differentObjects = append(differentObjects, dataSet.objects[i])
		}
	}

	sameObjectCountr := 0

	for i := 0; i < len(dataSet.objects); i++ {
		for j := 0; j < len(objects2); j++ {
			if dataSet.objects[i] == objects2[j] {
				sameObjectCountr++
			}
		}
	}

	if len(dataSet.objects) == len(objects2) && len(dataSet.objects) == sameObjectCountr {
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
