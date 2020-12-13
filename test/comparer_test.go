package test

import (
	"strconv"
	"testing"

	"kaanyilgin.com/dataComparer/model"
)

var testCases = []int{0, 1}

func TestCompareDifferentSize(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5}]`

	compareDataSet(firstSet, secondSet, false, "Two different sized jsons are not identical", t)
}

func TestCompareSameSizeDifferentInformation(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6}]`
	var secondSet = `[{"id":5},{"id":7}]`

	compareDataSet(firstSet, secondSet, false, "Datasets have different objects are not identical", t)
}

func TestCompareSameJsonRandomOrder(t *testing.T) {
	var firstSet = `[{"id":6},{"id":5}]`
	var secondSet = `[{"id":5},{"id":6}]`

	compareDataSet(firstSet, secondSet, true, "Datasets have same objects in different order are identical", t)
}

func TestCompareDifferentJsonWithDuplicatedValue(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5,"name":"Due"},{"id":5,"name":"Due"}]`

	compareDataSet(firstSet, secondSet, false, "Different datasets when one of them has duplicated matched objects are not identical", t)
}

func TestCompareDifferentJsonWithSecondDataSetHasDuplicatedExtraValue(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5,"name":"Due"},{"id":5,"name":"Due"},{"id":6,"name":"John"}]`

	compareDataSet(firstSet, secondSet, false, "Different datasets when one of them has one extra duplicated matched objects are not identical", t)
}

func TestCompareSameJsonWithOneDataSetHasBlankBetweenObjects(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5, "name":"Due"}, {"id":6, "name":"John"}]`

	compareDataSet(firstSet, secondSet, true, "Datasets are identical if one of them has some whitespace between attributes and objects", t)
}

func TestCompareDifferentJsonWithOneDataSetHasBlankInAttiribute(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":5, "name":"Due "}, {"id":6, "name":"John"}]`

	compareDataSet(firstSet, secondSet, false, "Datasets are different with when there is a different whitespace in the value", t)
}

func TestCompareDifferentJsonWithIdenticalAttributesInDifferentObjects(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"id":6,"name":"Due"},{"id":5,"name":"John"}]`

	compareDataSet(firstSet, secondSet, false, "Datasets are different but there are identical attributes in different objects", t)
}

func TestCompareSameJsonWithRandomAttributeOrder(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"name":"Due","id":5},{"name":"John","id":6}]`

	compareDataSet(firstSet, secondSet, true, "Datasets are same but object and attribute orders are random", t)
}

func TestCompareSameJsonWithRandomAttributeOrderWith3Items(t *testing.T) {
	var firstSet = `[{"id":6,"name":"John","extraProperty":true},{"id":6,"name":"John"},{"id":5,"name":"Due"}]`
	var secondSet = `[{"name":"Due","id":5},{"name":"John","id":6},{"name":"John","extraProperty":true,"id":6}]`

	compareDataSet(firstSet, secondSet, true, "Datasets are same but object and attribute orders are random with 3 items", t)
}

func TestCompareDifferentJsonSecondHasDifferentObject(t *testing.T) {
	var firstSet = `[{"id":5},{"id":6},{"id":5}]`
	var secondSet = `[{"id":5},{"id":7},{"id":6}]`

	compareDataSet(firstSet, secondSet, false, "Datasets are different when second has a different object", t)
}

func compareDataSet(dataset1 string, dataset2 string, expectedValue bool, errorMessage string, t *testing.T) {
	for i := 0; i < len(testCases); i++ {
		dataSet := createDataSet(i, dataset1)
		secondDataSet := createDataSet(i, dataset2)
		isEqual, _ := dataSet.IsEqual(secondDataSet)

		if isEqual != expectedValue {
			t.Errorf("TestCase: " + strconv.Itoa(i) + " failed. " + errorMessage)
		}
	}
}

func createDataSet(testCase int, data string) *model.DataSet {
	var dataSetComparer model.DataSetComparer
	var jsonParser model.JSONParser
	var dataSet *model.DataSet

	if testCase == 0 {
		dataSetComparer = model.LoopingTwoDataSetComparer{
			ObjectComparer: model.FindAttributeByKeyObjectComparer{},
		}
		jsonParser = &model.JSONObjectParser{}
		objects := jsonParser.ParseJSON(data).([]map[string]interface{})
		dataSet = model.InitDataSet(objects, dataSetComparer)
	} else {
		dataSetComparer = model.HashMapDataSetComparer{}

		jsonParser = &model.JSONObjectMapParser{}

		objects := jsonParser.ParseJSON(data).(model.JSONObjectMap)
		dataSet = model.InitDataSetHashMapObject(objects, dataSetComparer)
	}

	return dataSet
}
