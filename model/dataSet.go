package model

// DataSet is using for storing data
type DataSet struct {
	Data string
}

// Init creates a new DataSet object
func (dataSet DataSet) Init(data string) *DataSet {
	newDataSet := new(DataSet)
	newDataSet.Data = data
	return newDataSet
}

// Compare compares given dataset
func (dataSet DataSet) Compare(data string) bool {
	return len(dataSet.Data) == len(data)
}
