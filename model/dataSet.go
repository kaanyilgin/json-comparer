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

	return dataSet.compareDataSets(comparedData)
}

func (dataSet DataSet) compareDataSets(comparedDataSet *DataSet) bool {
	differentObjectCount := 0

	for _, object := range dataSet.objects {
		for _, objectCompared := range comparedDataSet.objects {
			if object.Compare(objectCompared) == false {
				differentObjectCount++
			} else {
				differentObjectCount--
			}
		}

		if dataSet.getObjectCount() == differentObjectCount {
			return false
		}
	}

	return 0 == differentObjectCount
}

func (dataSet DataSet) getObjectCount() int {
	return len(dataSet.objects)
}
