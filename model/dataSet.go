package model

import (
	"strings"
)

// DataSet is using for storing data
type DataSet struct {
	Data string
}

// NewDataSet creates a new DataSet object
func NewDataSet(data string) *DataSet {
	return &DataSet{
		Data: data,
	}
}

// Compare compares given dataset
func (dataSet DataSet) Compare(data string) bool {
	objects1 := parseJSON(dataSet.Data)
	objects2 := parseJSON(data)
	isSameSize := len(objects1) == len(objects2)

	if isSameSize == false {
		return false
	}

	differentObjects := make([]string, 0)

	for i := 0; i < len(objects1); i++ {
		if objects1[i] != objects2[i] {
			differentObjects = append(differentObjects, objects1[i])
		}
	}

	sameObjectCountr := 0

	for i := 0; i < len(objects1); i++ {
		for j := 0; j < len(objects2); j++ {
			if objects1[i] == objects2[j] {
				sameObjectCountr++
			}
		}
	}

	if len(objects1) == len(objects2) && len(objects1) == sameObjectCountr {
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
