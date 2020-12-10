package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

func Benchmark(b *testing.B) {
	reader := &infrastructure.FileReader{}
	dataSetComparer := new(model.LoopingTwoDataSetComparer)
	objectComparer := model.LoopingTwoObjectAttributeComparer{}
	dataset := model.NewDataSet("MOCK_DATA.json", reader, dataSetComparer, objectComparer)
	dataset2 := model.NewDataSet("MOCK_DATA.json", reader, dataSetComparer, objectComparer)

	isEqual, _ := dataset.Compare(dataset2)

	print(isEqual)
}
