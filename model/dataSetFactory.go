package model

const (
	DataSetArrayObjectType = iota
)

type DataSetFactory interface {
	CreateDataSet(dataSetType int, dataSetObjects interface{}) *DataSet
}

type DefaultDataSetFactory struct {
	dataSetComparer DataSetComparer
}

func InitDefaultDataSetFactory(dataSetComparer DataSetComparer) *DefaultDataSetFactory {
	return &DefaultDataSetFactory{
		dataSetComparer,
	}
}

func (d DefaultDataSetFactory) CreateDataSet(dataSetType int, dataSetObjects interface{}) *DataSet {
	if dataSetType == DataSetArrayObjectType {
		return InitDataSet(dataSetObjects.([]map[string]interface{}), d.dataSetComparer)
	}

	return nil
}
