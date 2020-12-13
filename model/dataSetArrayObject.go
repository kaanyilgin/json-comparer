package model

// DataSetArrayObject stored json object as array
type DataSet struct {
	objects         interface{}
	dataSetComparer DataSetComparer
}

// InitDataSetArrayObject creates a new DataSetArrayObject object
func InitDataSetArrayObject(objects []map[string]interface{}, dataSetComparer DataSetComparer) *DataSet {
	jsonObjects := make([]JSONObject, 0)

	for _, v := range objects {
		jsonObject := InitJSONObject(v)
		jsonObjects = append(jsonObjects, *jsonObject)
	}

	return &DataSet{
		dataSetComparer: dataSetComparer,
		objects:         jsonObjects,
	}
}

// InitDataSetArrayObject creates a new DataSetArrayObject object
func InitDataSetHashMapObject(objects []map[string]interface{}, dataSetComparer DataSetComparer) *DataSet {
	jsonObjects := make([]JSONObject, 0)

	for _, v := range objects {
		jsonObject := InitJSONObject(v)
		jsonObjects = append(jsonObjects, *jsonObject)
	}

	return &DataSet{
		dataSetComparer: dataSetComparer,
		objects:         jsonObjects,
	}
}

// NewDataSet creates a new DataSet object for testing
func NewDataSetTesting(objects JSONObjectMap, dataSetComparer DataSetComparer) *DataSet {
	return &DataSet{
		dataSetComparer: dataSetComparer,
		objects:         objects,
	}
}

func (dataSet *DataSet) IsEqual(comparedData interface{}) (bool, error) {
	dataSetType := comparedData.(*DataSet)
	isSameSize := dataSet.getObjectCount() == dataSetType.getObjectCount()

	if isSameSize == false {
		return false, nil
	}

	return dataSet.dataSetComparer.Compare(dataSet, dataSetType)
}

func (dataSet *DataSet) getObjectCount() int {
	return 1000 // TODO
}
