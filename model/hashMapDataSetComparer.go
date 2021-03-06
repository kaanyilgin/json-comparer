package model

//HashMapDataSetComparer compares two dataset creating hash of the objects
type HashMapDataSetComparer struct {
}

//Compare compares two dataset
func (h HashMapDataSetComparer) Compare(firstDataSet *DataSet, secondDataSet *DataSet) bool {
	firstDataSetObjects := firstDataSet.objects.(JSONObjectMap)
	secondDataSetObjects := secondDataSet.objects.(JSONObjectMap)

	for key, object := range firstDataSetObjects.dictionary {
		comparedObject := secondDataSetObjects.dictionary[key]

		if comparedObject == nil {
			return false
		}

		if comparedObject.jsonObjectCount != object.jsonObjectCount {
			return false
		}
	}

	return true
}
