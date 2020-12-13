package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/application"
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

var (
	reader         = &infrastructure.FileReader{}
	dataSetFactory = &model.DefaultDataSetFactory{}
)

func BenchmarkWithDirectlyFindAttributeKeyByMapKey(b *testing.B) {
	dataSetComparer := &model.LoopingTwoDataSetComparer{
		ObjectComparer: &model.FindAttributeByKeyObjectComparer{},
	}
	jsonParser := &model.JSONObjectParser{}
	dataSetCompare := &application.DataSetCompare{
		DataReader:     reader,
		JsonParser:     jsonParser,
		DataSetFactory: model.InitDefaultDataSetFactory(dataSetComparer),
	}

	isEqual, _ := dataSetCompare.CompareDataSets("MOCK_DATA.json", "MOCK_DATA.json", 0)

	print(isEqual)
}

func BenchmarkWithDirectlyFindObjectWithHashMap(b *testing.B) {
	dataSetComparer := &model.HashMapDataSetComparer{}
	jsonParser := &model.JSONObjectMapParser{}
	dataSetCompare := &application.DataSetCompare{
		DataReader:     reader,
		JsonParser:     jsonParser,
		DataSetFactory: model.InitDefaultDataSetFactory(dataSetComparer),
	}

	isEqual, _ := dataSetCompare.CompareDataSets("MOCK_DATA.json", "MOCK_DATA.json", 1)

	print(isEqual)
}
