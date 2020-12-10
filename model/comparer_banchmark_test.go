package model

import (
	"testing"

	"kaanyilgin.com/dataComparer/infrastructure"
)

var (
	dataSetComparer = new(LoopingTwoDataSetComparer)
	objectComparer  = LoopingTwoObjectAttributeComparer{}
)

func BenchmarkWithoutAttributeSorting(b *testing.B) {
	reader := &infrastructure.FileReader{}
	jsonParser := newSystemJSONParser()
	dataset := NewDataSet("MOCK_DATA.json", reader, dataSetComparer, objectComparer, jsonParser)
	dataset2 := NewDataSet("MOCK_DATA.json", reader, dataSetComparer, objectComparer, jsonParser)

	isEqual, _ := dataset.Compare(dataset2)

	print(isEqual)
}
