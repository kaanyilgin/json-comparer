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
	for i, object := range dataSet.objects {
		differentObjectCount := 0

		for j, objectCompared := range comparedDataSet.objects {
			if object.Compare(objectCompared) == false {
				differentObjectCount++
			} else {
				dataSet.objects = remove(dataSet.objects, i)
				comparedDataSet.objects = remove(comparedDataSet.objects, j)
				break
			}
		}

		objectIsExistInComparedDataSet := dataSet.getObjectCount() != differentObjectCount

		if objectIsExistInComparedDataSet == false {
			return false
		}
	}

	return true
}

func (dataSet DataSet) getObjectCount() int {
	return len(dataSet.objects)
}

func remove(slice []JSONObject, s int) []JSONObject {
	return append(slice[:s], slice[s+1:]...)
}
