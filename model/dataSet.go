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

	values1 := make([]byte, 0)
	values2 := make([]byte, 0)

	for i := 0; i < len(dataSet.Data); i++ {
		if dataSet.Data[i] != data[i] {
			values1 = append(values1, dataSet.Data[i])
			values2 = append(values2, data[i])
		}
	}

	sameByteCount := 0

	for i := 0; i < len(values1); i++ {
		for j := 0; j < len(values2); j++ {
			if values1[i] == values2[j] {
				sameByteCount++
			}
		}
	}

	if len(values1) == len(values2) && len(values1) == sameByteCount {
		return true
	}

	return false
}

func removeElement(values []byte, i int) {
	copy(values[i:], values[i+1:])
	emptyString := ""
	values[len(values)-1] = byte(emptyString[0])
	values = values[:len(values)-1]
}
