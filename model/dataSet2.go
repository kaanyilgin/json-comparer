package model

import "kaanyilgin.com/dataComparer/infrastructure"

type DataSet2 struct {
	objects         JSONObjectMap
	dataReader      infrastructure.DataReader
	dataSetComparer DataSetComparer
	objectComparer  ObjectComparer
	fileName        string
	jsonParser      JSONParser
}

func NewDataSet2(fileName string, dataReader infrastructure.DataReader, dataSetComparer DataSetComparer, objectComparer ObjectComparer, jsonParser JSONParser) *DataSet2 {
	return &DataSet2{
		dataReader:      dataReader,
		fileName:        fileName,
		dataSetComparer: dataSetComparer,
		objectComparer:  objectComparer,
		jsonParser:      jsonParser,
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
	return len(dataSet.getObjects().dictionary)
}

func (dataSet *DataSet2) readDataFromSource() JSONObjectMap {
	data, _ := dataSet.dataReader.Read(dataSet.fileName)
	serializedJSONObjects := dataSet.jsonParser.ParseJSON(data)
	return serializedJSONObjects
}

func (dataSet *DataSet2) getObjects() JSONObjectMap {
	if dataSet.objects.dictionary == nil {
		dataSet.objects = dataSet.readDataFromSource()
	}

	return dataSet.objects
}
