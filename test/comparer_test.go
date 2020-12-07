package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/model"
)

func TestCompareDifferentSize(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5}]`

	dataSet := model.NewDataSet(firstSet)
	isEqual := dataSet.Compare(secondSet)

	if isEqual != false {
		t.Errorf("Is eqaul was incorrect, two different sized dataset.")
	}
}

func TestCompareSameSizeDifferentInformation(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5},{"id":7}]`

	dataSet := model.NewDataSet(firstSet)
	isEqual := dataSet.Compare(secondSet)

	if isEqual != false {
		t.Errorf("Is eqaul was incorrect, two different information dataset.")
	}
}
func TestCompareSameJsonRandomOrder(t *testing.T) {
	var firstSet = `[{"id":6},{"id":5}]`
	var secondSet = `[{"id":5},{"id":7}]`

	dataSet := model.NewDataSet(firstSet)
	isEqual := dataSet.Compare(secondSet)

	if isEqual != true {
		t.Errorf("Is eqaul was incorrect, same json random order")
	}
}
