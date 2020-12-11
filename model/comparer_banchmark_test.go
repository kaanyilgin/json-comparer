package model

import (
	"testing"

	"kaanyilgin.com/dataComparer/infrastructure"
)

var (
	dataSetComparer = new(LoopingTwoDataSetComparer)
	reader          = &infrastructure.FileReader{}
	jsonParser      = newSystemJSONParser()
)

func BenchmarkWithoutAttributeSorting(b *testing.B) {
	objectComparer := &LoopingTwoObjectAttributeComparer{}
	mockReader1 := &mockDataReader{
		data: `[{"last_name":"Dumbreck","ip_address":"206.54.46.237","id":1,"gender":"Male","first_name":"Brannon","email":"bdumbreck0@flavors.me","country_code":"Serbia","country":"Ecuador","city":"Babahoyo","animal_name":"Indian leopard"}]`,
	}
	mockReader2 := &mockDataReader{
		data: `[{"id":1,"first_name":"Brannon","last_name":"Dumbreck","email":"bdumbreck0@flavors.me","gender":"Male","ip_address":"206.54.46.237","animal_name":"Indian leopard","city":"Babahoyo","country":"Ecuador","country_code":"Serbia"}]`,
	}
	dataset := NewDataSet("", mockReader1, dataSetComparer, objectComparer, jsonParser)
	dataset2 := NewDataSet("", mockReader2, dataSetComparer, objectComparer, jsonParser)

	isEqual, _ := dataset.Compare(dataset2)

	print(isEqual)
}

func BenchmarkWithDirectlyFindAttributeKeyByMapKey(b *testing.B) {
	objectComparer := &FindAttributeByKeyObjectComparer{}
	mockReader1 := &mockDataReader{
		data: `[{"last_name":"Dumbreck","ip_address":"206.54.46.237","id":1,"gender":"Male","first_name":"Brannon","email":"bdumbreck0@flavors.me","country_code":"Serbia","country":"Ecuador","city":"Babahoyo","animal_name":"Indian leopard"}]`,
	}
	mockReader2 := &mockDataReader{
		data: `[{"id":1,"first_name":"Brannon","last_name":"Dumbreck","email":"bdumbreck0@flavors.me","gender":"Male","ip_address":"206.54.46.237","animal_name":"Indian leopard","city":"Babahoyo","country":"Ecuador","country_code":"Serbia"}]`,
	}

	dataset := NewDataSet("", mockReader1, dataSetComparer, objectComparer, jsonParser)
	dataset2 := NewDataSet("", mockReader2, dataSetComparer, objectComparer, jsonParser)

	isEqual, _ := dataset.Compare(dataset2)

	print(isEqual)
}
