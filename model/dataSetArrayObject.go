package model

// DataSetArrayObject stored json object as array
type DataSetArrayObject struct {
	objects         []JSONObject
	dataSetComparer DataSetComparer
}

// NewDataSet creates a new DataSet object
func NewDataSet(dataSetComparer DataSetComparer) *DataSetArrayObject {
	return &DataSetArrayObject{
		dataSetComparer: dataSetComparer,
	}
}

// NewDataSet creates a new DataSet object for testing
func NewDataSetTesting(objects []JSONObject, dataSetComparer DataSetComparer) *DataSetArrayObject {
	return &DataSetArrayObject{
		dataSetComparer: dataSetComparer,
		objects:         objects,
	}
}

func (dataSet *DataSetArrayObject) IsEqual(comparedData interface{}) (bool, error) {
	dataSetType := comparedData.(*DataSetArrayObject)
	isSameSize := dataSet.getObjectCount() == dataSetType.getObjectCount()

	if isSameSize == false {
		return false, nil
	}

	return dataSet.dataSetComparer.Compare(dataSet, dataSetType)
}

func (dataSet *DataSetArrayObject) getObjectCount() int {
	return len(dataSet.objects)
}
