package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/application"
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

var (
	reader = &infrastructure.FileReader{}
)

func BenchmarkWithDirectlyFindAttributeKeyByMapKey(b *testing.B) {
	dataSetComparer := &model.LoopingTwoDataSetComparer{
		ObjectComparer: &model.FindAttributeByKeyObjectComparer{},
	}
	jsonParser := &model.JSONObjectParser{}
	dataSetCompare := &application.DataSetCompare{
		DataReader:     reader,
		JSONParser:     jsonParser,
		DataSetFactory: model.InitDefaultDataSetFactory(dataSetComparer),
	}

	dataSetCompare.CompareDataSets("MOCK_DATA.json", "MOCK_DATA_shuffled.json", 0)
}

func BenchmarkWithDirectlyFindObjectWithHashMap(b *testing.B) {
	dataSetComparer := &model.HashMapDataSetComparer{}
	jsonParser := &model.JSONObjectMapParser{}
	dataSetCompare := &application.DataSetCompare{
		DataReader:     reader,
		JSONParser:     jsonParser,
		DataSetFactory: model.InitDefaultDataSetFactory(dataSetComparer),
	}

	dataSetCompare.CompareDataSets("MOCK_DATA.json", "MOCK_DATA_shuffled.json", 1)
}
