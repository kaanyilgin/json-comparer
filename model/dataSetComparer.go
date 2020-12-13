package model

// DataSetComparer provides an interface to compare two dataset
type DataSetComparer interface {
	Compare(firstDataSet *DataSet, secondDataSet *DataSet) (bool, error)
}

//LoopingTwoDataSetComparer compares two dataset going through whole dataset
type LoopingTwoDataSetComparer struct {
	ObjectComparer ObjectComparer
}

//Compare compares two dataset
func (l LoopingTwoDataSetComparer) Compare(firstDataSet *DataSet, secondDataSet *DataSet) (bool, error) {
	firstDataSetObjects := firstDataSet.objects.([]JSONObject)
	secondDataSetObjects := secondDataSet.objects.([]JSONObject)

	for i := 0; i < len(firstDataSetObjects); i++ {
		differentObjectCount := 0
		object := firstDataSetObjects[i]

		for j := 0; j < len(secondDataSetObjects); j++ {
			objectCompared := secondDataSetObjects[j]

			if isEqual, _ := l.ObjectComparer.Compare(object, objectCompared); isEqual == false {
				differentObjectCount++
			} else {
				firstDataSetObjects = remove(firstDataSetObjects, i)
				secondDataSetObjects = remove(secondDataSetObjects, j)
				differentObjectCount--
				i--
				break
			}

			objectIsExistInComparedDataSet := len(firstDataSetObjects) != differentObjectCount

			if objectIsExistInComparedDataSet == false {
				return false, nil
			}
		}
	}

	return true, nil
}

func remove(slice []JSONObject, s int) []JSONObject {
	return append(slice[:s], slice[s+1:]...)
}

//LoopingTwoDataSetComparer compares two dataset going through whole dataset
type HashMapDataSetComparer struct {
}

//Compare compares two dataset
func (h HashMapDataSetComparer) Compare(firstDataSet *DataSet, secondDataSet *DataSet) (bool, error) {
	firstDataSetObjects := firstDataSet.objects.(JSONObjectMap)
	secondDataSetObjects := secondDataSet.objects.(JSONObjectMap)

	for key := range firstDataSetObjects.dictionary {
		comparedObject := secondDataSetObjects.dictionary[key]

		if comparedObject == nil {
			return false, nil
		}
	}

	return true, nil
}
