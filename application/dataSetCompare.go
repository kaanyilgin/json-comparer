package application

import (
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

type DataSetCompare struct {
	DataReader infrastructure.DataReader
	JsonParser model.JSONParser
}

func (d DataSetCompare) CompareDataSets(fileName1 string, fileName2 string) (bool, error) {
	firstDataSet := d.readDataFromSource(fileName1).(model.IDataSet)
	secondDataSet := d.readDataFromSource(fileName2).(model.IDataSet)

	return firstDataSet.IsEqual(secondDataSet)
}

func (d DataSetCompare) readDataFromSource(fileName string) interface{} {
	data, _ := d.DataReader.Read(fileName)
	serializedJSONObjects := d.JsonParser.ParseJSON(data)
	return serializedJSONObjects
}
