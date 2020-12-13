package model

// DataSet stored json object as array
type DataSet struct {
	objects         IJSONObject
	dataSetComparer DataSetComparer
}

// InitDataSet creates a new DataSet object
func InitDataSet(objects []map[string]interface{}, dataSetComparer DataSetComparer) *DataSet {
	jsonObjects := make([]JSONObject, 0)

	for _, v := range objects {
		jsonObject := InitJSONObject(v)
		jsonObjects = append(jsonObjects, *jsonObject)
	}

	return &DataSet{
		dataSetComparer: dataSetComparer,
		objects: &JSONObjectArray{
			dictonary: jsonObjects,
		},
	}
}

// InitDataSetHashMapObject creates a new DataSet object
func InitDataSetHashMapObject(objects JSONObjectMap, dataSetComparer DataSetComparer) *DataSet {
	return &DataSet{
		objects,
		dataSetComparer,
	}
}

// IsEqual compares the given data set
func (dataSet *DataSet) IsEqual(comparedData *DataSet) (bool, error) {
	isSameSize := dataSet.objects.GetLength() == comparedData.objects.GetLength()

	if isSameSize == false {
		return false, nil
	}

	return dataSet.dataSetComparer.Compare(dataSet, comparedData)
}
