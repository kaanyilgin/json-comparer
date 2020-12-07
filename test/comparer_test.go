package test

import (
	"testing"

	"kaanyilgin.com/dataComparer/model"
)

func TestCompareDifferenSizeJSONFiles(t *testing.T) {
	var firstSet = "[{id:5},{id:6}]"
	var secondSet = "[{id:5}]"

	dataSet := new(model.DataSet)
	dataSet.Init(firstSet)
	isEqual := dataSet.Compare(secondSet)

	if isEqual != false {
		t.Errorf("Is eqaul was incorrect, two different sized json files.")
	}
}