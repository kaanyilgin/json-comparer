package model

// DataSet is using for storing data
type DataSet struct {
	Data string
}

// NewDataSet creates a new DataSet object
func NewDataSet(data string) *DataSet {
	return &DataSet{
		Data: data,
	}
}

// Compare compares given dataset
func (dataSet DataSet) Compare(data string) bool {
	isSameSize := len(dataSet.Data) == len(data)

	if isSameSize == false {
		return false
	}

	for i := 0; i < len(dataSet.Data); i++ {
		if dataSet.Data[i] != data[i] {
			return false
		}
	}

	return true
}
