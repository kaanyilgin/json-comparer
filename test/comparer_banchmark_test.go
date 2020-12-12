package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/application"
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

var (
	dataSetComparer = new(model.LoopingTwoDataSetComparer)
	reader          = &infrastructure.FileReader{}
	jsonParser      = &model.JSONObjectParser{}
)

func BenchmarkWithDirectlyFindAttributeKeyByMapKey(b *testing.B) {
	dataSetCompare := &application.DataSetCompare{
		DataReader: reader,
		JsonParser: jsonParser,
	}

	isEqual, _ := dataSetCompare.CompareDataSets("MOCK_DATA.json", "MOCK_DATA.json")

	print(isEqual)
}

// func BenchmarkObjectHashMap(b *testing.B) {
// 	objectComparer := &FindAttributeByKeyObjectComparer{}

// 	dataset := NewDataSet2("MOCK_DATA.json", reader, dataSetComparer, objectComparer)
// 	dataset2 := NewDataSet2("MOCK_DATA.json", reader, dataSetComparer, objectComparer)

// 	isEqual, _ := dataset.IsEqual(dataset2)

// 	print(isEqual)
// }
