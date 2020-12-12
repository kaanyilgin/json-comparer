package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"kaanyilgin.com/dataComparer/infrastructure"
)

type IDataSet interface{
	(dataSet interface[}]) IsEqual(comparedData *interface[}) (bool, error)
}

// DataSet is using for storing data
type DataSet struct {
	objects         []JSONObject
	dataReader      infrastructure.DataReader
	dataSetComparer DataSetComparer
	objectComparer  ObjectComparer
	fileName        string
	jsonParser      JSONParser
}

func (dataSet DataSet) IsEqual(comparedData *DataSet) (bool, error){
	return false, nil
}

// JSONObjectMap stands TODO
type JSONObjectMap struct {
	dictionary     map[string]*JSONObject
	objectComparer ObjectComparer
}

// UnmarshalJSON custom json UnmarshalJSON
func (j *JSONObjectMap) UnmarshalJSON(data []byte) error {

	var unMarshalledData []interface{}
	if err := json.Unmarshal(data, &unMarshalledData); err != nil {
		return err
	}

	j.dictionary = make(map[string]*JSONObject, 0)

	for _, object := range unMarshalledData {
		jsonObject2 := object.(map[string]interface{})
		hash := asSha256(jsonObject2)
		j.dictionary[hash] = NewJSONObject(jsonObject2, j.objectComparer)
	}

	return nil
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

type DataSet2 struct {
	dataSet	DataSet
	objects JSONObjectMap
}

// NewDataSet creates a new DataSet object
func NewDataSet(fileName string, dataReader infrastructure.DataReader, dataSetComparer DataSetComparer, objectComparer ObjectComparer, jsonParser JSONParser) *DataSet {
	return &DataSet{
		dataReader:      dataReader,
		fileName:        fileName,
		dataSetComparer: dataSetComparer,
		objectComparer:  objectComparer,
		jsonParser:      jsonParser,
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
	serializedJSONObjects := dataSet.jsonParser.ParseJSON(data)
	objects := make([]JSONObject, 0)

	for _, attributes := range serializedJSONObjects.dictionary {
		object := NewJSONObject(attributes.attributes, dataSet.objectComparer)
		objects = append(objects, *object)
	}

	return objects
}

func (dataSet *DataSet) getObjects() []JSONObject {
	if dataSet.objects == nil {
		dataSet.objects = dataSet.readDataFromSource()
	}

	return dataSet.objects
}
