package model

const (
	DataSetArrayObjectType = iota
)

type DataSetFactory interface {
	CreateDataSetFactory(dataSetType int, dataSetObjects interface{}) IDataSet
}

type DefaultDataSetFactory struct {
	dataSetComparer DataSetComparer
}

func InitDefaultDataSetFactory(dataSetComparer DataSetComparer) *DefaultDataSetFactory {
	return &DefaultDataSetFactory{
		dataSetComparer,
	}
}

func (d DefaultDataSetFactory) CreateDataSetFactory(dataSetType int, dataSetObjects interface{}) IDataSet {
	if dataSetType == DataSetArrayObjectType {
		return InitDataSetArrayObject(dataSetObjects.([]map[string]interface{}), d.dataSetComparer)
	}

	return nil
}
