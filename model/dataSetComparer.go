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
	firstDataSetObjects := firstDataSet.objects.(*JSONObjectArray)
	secondDataSetObjects := secondDataSet.objects.(*JSONObjectArray)

	for i := 0; i < firstDataSetObjects.GetLength(); i++ {
		differentObjectCount := 0
		object := firstDataSetObjects.dictonary[i]

		for j := 0; j < secondDataSetObjects.GetLength(); j++ {
			objectCompared := secondDataSetObjects.dictonary[j]

			if isEqual, _ := l.ObjectComparer.Compare(object, objectCompared); isEqual == false {
				differentObjectCount++
			} else {
				firstDataSetObjects.dictonary = remove(firstDataSetObjects.dictonary, i)
				secondDataSetObjects.dictonary = remove(secondDataSetObjects.dictonary, j)
				differentObjectCount--
				i--
				break
			}

			objectIsExistInComparedDataSet := firstDataSetObjects.GetLength() != differentObjectCount

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

	for key, object := range firstDataSetObjects.dictionary {
		comparedObject := secondDataSetObjects.dictionary[key]

		if comparedObject == nil {
			return false, nil
		} else {
			if comparedObject.jsonObjectCount != object.jsonObjectCount {
				return false, nil
			}
		}
	}

	return true, nil
}
