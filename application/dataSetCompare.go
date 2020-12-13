package application

import (
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

type DataSetCompare struct {
	DataReader     infrastructure.DataReader
	JsonParser     model.JSONParser
	DataSetFactory model.DataSetFactory
}

func (d DataSetCompare) CompareDataSets(fileName1 string, fileName2 string, dataSetType int) (bool, error) {
	firstDataSetSource, _ := d.readDataFromSource(fileName1)
	secondDataSetSource, _ := d.readDataFromSource(fileName2)

	firstDataSet, err := d.DataSetFactory.CreateDataSet(dataSetType, firstDataSetSource)

	if err != nil {
		return false, err
	}

	secondDataSet, err := d.DataSetFactory.CreateDataSet(dataSetType, secondDataSetSource)

	if err != nil {
		return false, err
	}

	return firstDataSet.IsEqual(secondDataSet)
}

func (d DataSetCompare) readDataFromSource(fileName string) (interface{}, error) {
	data, err := d.DataReader.Read(fileName)
	serializedJSONObjects := d.JsonParser.ParseJSON(data)
	return serializedJSONObjects, err
}
