package model

// DataSetComparer provides an interface to compare two dataset
type DataSetComparer interface {
	Compare(firstDataSet *DataSet, secondDataSet *DataSet) bool
}
