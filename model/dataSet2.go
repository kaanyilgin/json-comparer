package model

type DataSet2 struct {
	objects         JSONObjectMap
	dataSetComparer DataSetComparer
}

func NewDataSet2(dataSetComparer DataSetComparer) *DataSet2 {
	return &DataSet2{
		dataSetComparer: dataSetComparer,
	}
}

func (dataSet *DataSet2) IsEqual(comparedData interface{}) (bool, error) {
	dataSetType := comparedData.(*DataSet2)
	isSameSize := dataSet.getObjectCount() == dataSetType.getObjectCount()

	if isSameSize == false {
		return false, nil
	}

	comparer := new(HashMapDataSetComparer)
	return comparer.Compare(dataSet, dataSetType)
}

func (dataSet *DataSet2) getObjectCount() int {
	return len(dataSet.objects.dictionary)
}
