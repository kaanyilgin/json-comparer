package model

import "errors"

const (
	// DataSetArrayObjectType is a type for LoopingTwoDataSetComparer
	DataSetArrayObjectType = iota
	// DataSetObjectHashKeyType is a type for HashMapDataSetComparer
	DataSetObjectHashKeyType = iota
)

// DataSetFactory creates dataset with the given object and comparing algorithm
type DataSetFactory interface {
	CreateDataSet(dataSetType int, dataSetObjects interface{}) (*DataSet, error)
}

// DefaultDataSetFactory creates dataset
type DefaultDataSetFactory struct {
	dataSetComparer DataSetComparer
}

// InitDefaultDataSetFactory is a constructor for DefaultDataSetFactory
func InitDefaultDataSetFactory(dataSetComparer DataSetComparer) *DefaultDataSetFactory {
	return &DefaultDataSetFactory{
		dataSetComparer,
	}
}

// CreateDataSet creates dataSet with the value
func (d DefaultDataSetFactory) CreateDataSet(dataSetType int, dataSetObjects interface{}) (*DataSet, error) {
	if dataSetType == DataSetArrayObjectType {
		return InitDataSet(dataSetObjects.([]map[string]interface{}), d.dataSetComparer), nil
	}
	if dataSetType == DataSetObjectHashKeyType {
		return InitDataSetHashMapObject(dataSetObjects.(JSONObjectMap), d.dataSetComparer), nil
	}

	return nil, errors.New("Undefined dataset type")
}
