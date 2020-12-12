package model

import "kaanyilgin.com/dataComparer/infrastructure"

// DataSetArrayObject stored json object as array
type DataSetArrayObject struct {
	objects         []JSONObject
	dataReader      infrastructure.DataReader
	dataSetComparer DataSetComparer
	objectComparer  ObjectComparer
	fileName        string
	jsonParser      JSONParser
}

// NewDataSet creates a new DataSet object
func NewDataSet(fileName string, dataReader infrastructure.DataReader, dataSetComparer DataSetComparer, objectComparer ObjectComparer, jsonParser JSONParser) *DataSetArrayObject {
	return &DataSetArrayObject{
		dataReader:      dataReader,
		fileName:        fileName,
		dataSetComparer: dataSetComparer,
		objectComparer:  objectComparer,
		jsonParser:      jsonParser,
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
	return len(dataSet.getObjects())
}

func (dataSet *DataSetArrayObject) readDataFromSource() []JSONObject {
	data, _ := dataSet.dataReader.Read(dataSet.fileName)
	serializedJSONObjects := dataSet.jsonParser.ParseJSON(data)
	objects := make([]JSONObject, 0)

	for _, attributes := range serializedJSONObjects.dictionary {
		object := NewJSONObject(attributes.attributes, dataSet.objectComparer)
		objects = append(objects, *object)
	}

	return objects
}

func (dataSet *DataSetArrayObject) getObjects() []JSONObject {
	if dataSet.objects == nil {
		dataSet.objects = dataSet.readDataFromSource()
	}

	return dataSet.objects
}
