package application

import (
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

// DataSetCompare compares two json files
type DataSetCompare struct {
	DataReader     infrastructure.DataReader
	JSONParser     model.JSONParser
	DataSetFactory model.DataSetFactory
}

// CompareDataSets reads the files from the sources and compares them
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

	return firstDataSet.IsEqual(secondDataSet), nil
}

func (d DataSetCompare) readDataFromSource(fileName string) (interface{}, error) {
	data, err := d.DataReader.Read(fileName)
	serializedJSONObjects := d.JSONParser.ParseJSON(data)
	return serializedJSONObjects, err
}
