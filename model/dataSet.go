package model

type IDataSet interface {
	IsEqual(comparedData interface{}) (bool, error)
}
