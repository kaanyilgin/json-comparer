package model

import "errors"

const (
	DataSetArrayObjectType   = iota
	DataSetObjectHashKeyType = iota
)

type DataSetFactory interface {
	CreateDataSet(dataSetType int, dataSetObjects interface{}) (*DataSet, error)
}

type DefaultDataSetFactory struct {
	dataSetComparer DataSetComparer
}

func InitDefaultDataSetFactory(dataSetComparer DataSetComparer) *DefaultDataSetFactory {
	return &DefaultDataSetFactory{
		dataSetComparer,
	}
}

func (d DefaultDataSetFactory) CreateDataSet(dataSetType int, dataSetObjects interface{}) (*DataSet, error) {
	if dataSetType == DataSetArrayObjectType {
		return InitDataSet(dataSetObjects.([]map[string]interface{}), d.dataSetComparer), nil
	}
	if dataSetType == DataSetObjectHashKeyType {
		return InitDataSetHashMapObject(dataSetObjects.(JSONObjectMap), d.dataSetComparer), nil
	}

	return nil, errors.New("Undefined dataset type")
}
