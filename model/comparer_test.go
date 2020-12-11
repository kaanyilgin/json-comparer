package model

import (
	"testing"
)

func TestCompareDifferentSize(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != false {
		t.Errorf("Two different sized jsons are not identical")
	}
}

func TestCompareSameSizeDifferentInformation(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5},{"id":7}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != false {
		t.Errorf("Datasets have different objects are not identical")
	}
}

func TestCompareSameJsonRandomOrder(t *testing.T) {
	var firstSet = `[{"id":6},{"id":5}]`
	var secondSet = `[{"id":5},{"id":6}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != true {
		t.Errorf("Datasets have same objects in different order are identical")
	}
}

func TestCompareDifferentJsonWithDuplicatedValue(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5,"name":"Due"},{"id":5,"name":"Due"}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != false {
		t.Errorf("Different datasets when on of them has duplicated matched objects are not identical")
	}
}

func TestCompareSameJsonWithOneDataSetHasBlankBetweenObjects(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5, "name":"Due"}, {"id":6, "name":"John"}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != true {
		t.Errorf("Datasets are identical if one of them has some whitespace between attributes and objects")
	}
}

func TestCompareDifferentJsonWithOneDataSetHasBlankInAttiribute(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5, "name":"Due "}, {"id":6, "name":"John"}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != false {
		t.Errorf("Datasets are different with when there is a different whitespace in the value")
	}
}

func TestCompareDifferentJsonWithIdenticalAttributesInDifferentObjects(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":6,"name":"Due"},{"id":5,"name":"John"}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != false {
		t.Errorf("Datasets are different but there are identical attributes in different objects")
	}
}

func TestCompareSameJsonWithRandomAttributeOrder(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"name":"Due","id":5},{"name":"John","id":6}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != true {
		t.Errorf("Datasets are same but object and attribute orders are random")
	}
}

func TestCompareSameJsonWithRandomAttributeOrderWith3Items(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John","extraProperty":true},{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"name":"Due","id":5},{"name":"John","id":6},{"name":"John","extraProperty":true,"id":6}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != true {
		t.Errorf("Datasets are same but object and attribute orders are random with 3 items")
	}
}

func TestCompareDifferentJsonSecondHasDifferentObject(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6},{"id":5}]`
	var secondSet = `[{"id":5},{"id":7},{"id":6}]`

	isEqual := compareDataSet(firstSet, secondSet)

	if isEqual != false {
		t.Errorf("Datasets are different when second has a different object")
	}
}

func compareDataSet(dataset1 string, dataset2 string) bool {
	dataSet := createDataSet(dataset1)
	secondDataSet := createDataSet(dataset2)
	isEqual, _ := dataSet.Compare(secondDataSet)
	return isEqual
}

func createDataSet(data string) *DataSet {
	mockDataReader := &mockDataReader{
		data: data,
	}
	dataSetComparer := LoopingTwoDataSetComparer{}
	objectComparer := FindAttributeByKeyObjectComparer{}
	jsonParser := &SystemJSONParser{}

	return NewDataSet("", mockDataReader, dataSetComparer, objectComparer, jsonParser)
}

type mockDataReader struct {
	DataReader,
	data string
}

//Read returns the file content
func (m mockDataReader) Read(fileName string) (string, error) {
	return m.data, nil
}
