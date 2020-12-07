package test

import "testing"

func TestCompareDifferenSizeJsonFiles(t *testing.T) {
	var firstSet = "[{id:5},{id:6}]"
	var secondSet = "[{id:5}]"

	dataSet := new(DataSet)
	dataSet.Init(firstSet)
	isEqual := dataSet.Compare(secondSet)

	if isEqaul != false {
		t.Errorf("Is eqaul was incorrect, two different sized json files.")
	}
}
