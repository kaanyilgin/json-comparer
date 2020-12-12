package model

// DataSetComparer provides an interface to compare two dataset
type DataSetComparer interface {
	Compare(firstDataSet *DataSetArrayObject, secondDataSet *DataSetArrayObject) (bool, error)
}

//LoopingTwoDataSetComparer compares two dataset going through whole dataset
type LoopingTwoDataSetComparer struct {
	ObjectComparer ObjectComparer
}

//Compare compares two dataset
func (l LoopingTwoDataSetComparer) Compare(firstDataSet *DataSetArrayObject, secondDataSet *DataSetArrayObject) (bool, error) {
	for i := 0; i < firstDataSet.getObjectCount(); i++ {
		differentObjectCount := 0
		object := firstDataSet.objects[i]

		for j := 0; j < secondDataSet.getObjectCount(); j++ {
			objectCompared := secondDataSet.objects[j]

			if isEqual, _ := object.Compare(l.ObjectComparer, objectCompared); isEqual == false {
				differentObjectCount++
			} else {
				firstDataSet.objects = remove(firstDataSet.objects, i)
				secondDataSet.objects = remove(secondDataSet.objects, j)
				differentObjectCount--
				i--
				break
			}

			objectIsExistInComparedDataSet := firstDataSet.getObjectCount() != differentObjectCount

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
func (h HashMapDataSetComparer) Compare(firstDataSet *DataSet2, secondDataSet *DataSet2) (bool, error) {
	for key := range firstDataSet.objects.dictionary {
		comparedObject := secondDataSet.objects.dictionary[key]

		if comparedObject == nil {
			return false, nil
		}
	}

	return true, nil
}
