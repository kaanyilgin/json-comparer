package model

import "kaanyilgin.com/dataComparer/infrastructure"

// DataSet is using for storing data
type DataSet struct {
	objects    []JSONObject
	dataReader infrastructure.DataReader
	fileName   string
}

// NewDataSet creates a new DataSet object
func NewDataSet(fileName string, dataReader infrastructure.DataReader) *DataSet {
	return &DataSet{
		dataReader: dataReader,
		fileName:   fileName,
	}
}

// Compare compares given dataset
func (dataSet DataSet) Compare(comparedData *DataSet) (bool, error) {
	isSameSize := dataSet.getObjectCount() == comparedData.getObjectCount()

	if isSameSize == false {
		return false, nil
	}

	return dataSet.compareDataSets(comparedData), nil
}

func (dataSet DataSet) compareDataSets(comparedDataSet *DataSet) bool {
	for i := 0; i < dataSet.getObjectCount(); i++ {
		differentObjectCount := 0
		object := dataSet.getObjects()[i]

		for j := 0; j < comparedDataSet.getObjectCount(); j++ {
			objectCompared := comparedDataSet.getObjects()[j]

			if object.Compare(objectCompared) == false {
				differentObjectCount++
			} else {
				dataSet.objects = remove(dataSet.getObjects(), i)
				comparedDataSet.objects = remove(comparedDataSet.getObjects(), j)
				differentObjectCount--
				i--
				break
			}

			objectIsExistInComparedDataSet := dataSet.getObjectCount() != differentObjectCount

			if objectIsExistInComparedDataSet == false {
				return false
			}
		}
	}

	return true
}

func (dataSet *DataSet) getObjectCount() int {
	return len(dataSet.getObjects())
}

func remove(slice []JSONObject, s int) []JSONObject {
	return append(slice[:s], slice[s+1:]...)
}

func (dataSet *DataSet) readDataFromSource() []JSONObject {
	data, _ := dataSet.dataReader.Read(dataSet.fileName)
	return ParseJSON(data)
}

func (dataSet *DataSet) getObjects() []JSONObject {
	if dataSet.objects == nil {
		dataSet.objects = dataSet.readDataFromSource()
	}

	return dataSet.objects
}
