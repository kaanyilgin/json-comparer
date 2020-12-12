package model

// DataSetArrayObject stored json object as array
type DataSetArrayObject struct {
	objects         []JSONObject
	dataSetComparer DataSetComparer
}

// InitDataSetArrayObject creates a new DataSetArrayObject object
func InitDataSetArrayObject(objects []map[string]interface{}, dataSetComparer DataSetComparer) *DataSetArrayObject {
	jsonObjects := make([]JSONObject, 0)

	for _, v := range objects {
		jsonObject := NewJSONObject(v)
		jsonObjects = append(jsonObjects, *jsonObject)
	}

	return &DataSetArrayObject{
		dataSetComparer: dataSetComparer,
		objects:         jsonObjects,
	}
}

// NewDataSet creates a new DataSet object for testing
func NewDataSetTesting(objects []map[string]interface{}, dataSetComparer DataSetComparer) *DataSetArrayObject {
	jsonObjects := make([]JSONObject, 0)

	for _, v := range objects {
		jsonObject := NewJSONObject(v)
		jsonObjects = append(jsonObjects, *jsonObject)
	}

	return &DataSetArrayObject{
		dataSetComparer: dataSetComparer,
		objects:         jsonObjects,
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
