package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/model"
)

func TestCompareDifferentSize(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != false {
		t.Errorf("Is eqaul was incorrect, two different sized dataset.")
	}
}

func TestCompareSameSizeDifferentInformation(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5},{"id":7}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != false {
		t.Errorf("Is eqaul was incorrect, two different information dataset.")
	}
}

func TestCompareSameJsonRandomOrder(t *testing.T) {
	var firstSet = `[{"id":6},{"id":5}]`
	var secondSet = `[{"id":5},{"id":6}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != true {
		t.Errorf("Is eqaul was incorrect, same json random order")
	}
}

func TestCompareDifferentJsonWithDuplicatedValue(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5,"name":"Due"},{"id":5,"name":"Due"}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != false {
		t.Errorf("Is eqaul was incorrect, different json with duplicated value")
	}
}

func TestCompareSameJsonWithOneDataSetHasBlankBetweenObjects(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5, "name":"Due"}, {"id":6, "name":"John"}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != true {
		t.Errorf("Datasets are identical but one of them has blank between objects and attributes.")
	}
}

func TestCompareDifferentJsonWithOneDataSetHasBlankInAttiribute(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5, "name":"Due "}, {"id":6, "name":"John"}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != false {
		t.Errorf("Datasets are different with one blank in attirubte")
	}
}

func TestCompareDifferentJsonWithIdenticalAttributesInDifferentObjects(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":6,"name":"Due"},{"id":5,"name":"John"}]`

	dataSet := model.NewDataSet(firstSet)
	secondDataSet := model.NewDataSet(secondSet)
	isEqual := dataSet.Compare(secondDataSet)

	if isEqual != false {
		t.Errorf("Datasets are different but there are identical attributes in different objects")
	}
}
