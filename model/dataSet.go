package model

// DataSet is using for storing data
type DataSet struct {
	Data string
}

// Init creates a new DataSet object
func (p DataSet) Init(data string) *DataSet {
	dataSet := new(DataSet)
	dataSet.Data = data
	return dataSet
}

// Compare compares given dataset
func (p DataSet) Compare(data string) bool {
	return true
}
