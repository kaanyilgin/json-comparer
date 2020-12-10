package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

func Benchmark(b *testing.B) {
	reader := &infrastructure.FileReader{}
	objectComparer := new(model.LoopingTwoDataSetComparer)
	dataset := model.NewDataSet("MOCK_DATA.json", reader, objectComparer)
	dataset2 := model.NewDataSet("MOCK_DATA.json", reader, objectComparer)

	isEqual, _ := dataset.Compare(dataset2)

	print(isEqual)
}
