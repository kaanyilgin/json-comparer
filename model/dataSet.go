package model

import "kaanyilgin.com/dataComparer/infrastructure"

// DataSet is using for storing data
type DataSet struct {
	objects         []JSONObject
	dataReader      infrastructure.DataReader
	dataSetComparer DataSetComparer
	objectComparer  ObjectComparer
	fileName        string
}

// NewDataSet creates a new DataSet object
func NewDataSet(fileName string, dataReader infrastructure.DataReader, dataSetComparer DataSetComparer, objectComparer ObjectComparer) *DataSet {
	return &DataSet{
		dataReader:      dataReader,
		fileName:        fileName,
		dataSetComparer: dataSetComparer,
		objectComparer:  objectComparer,
	}
}

// Compare compares given dataset
func (dataSet DataSet) Compare(comparedData *DataSet) (bool, error) {
	isSameSize := dataSet.getObjectCount() == comparedData.getObjectCount()

	if isSameSize == false {
		return false, nil
	}

	return dataSet.dataSetComparer.Compare(&dataSet, comparedData)
}

func (dataSet *DataSet) getObjectCount() int {
	return len(dataSet.getObjects())
}

func (dataSet *DataSet) readDataFromSource() []JSONObject {
	data, _ := dataSet.dataReader.Read(dataSet.fileName)
	return ParseJSON(data, dataSet.objectComparer)
}

func (dataSet *DataSet) getObjects() []JSONObject {
	if dataSet.objects == nil {
		dataSet.objects = dataSet.readDataFromSource()
	}

	return dataSet.objects
}
