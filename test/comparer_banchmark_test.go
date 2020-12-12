package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/application"
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

var (
	dataSetComparer = &model.LoopingTwoDataSetComparer{
		ObjectComparer: &model.FindAttributeByKeyObjectComparer{},
	}
	reader         = &infrastructure.FileReader{}
	jsonParser     = &model.JSONObjectParser{}
	dataSetFactory = &model.DefaultDataSetFactory{}
)

func BenchmarkWithDirectlyFindAttributeKeyByMapKey(b *testing.B) {
	dataSetCompare := &application.DataSetCompare{
		DataReader:     reader,
		JsonParser:     jsonParser,
		DataSetFactory: model.InitDefaultDataSetFactory(dataSetComparer),
	}

	isEqual, _ := dataSetCompare.CompareDataSets("MOCK_DATA.json", "MOCK_DATA.json")

	print(isEqual)
}
