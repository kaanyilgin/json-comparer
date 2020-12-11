package model

// DataSetComparer provides an interface to compare two dataset
type DataSetComparer interface {
	Compare(firstDataSet *DataSet, secondDataSet *DataSet) (bool, error)
}

//LoopingTwoDataSetComparer compares two dataset going through whole dataset
type LoopingTwoDataSetComparer struct {
}

//Compare compares two dataset
func (l LoopingTwoDataSetComparer) Compare(firstDataSet *DataSet, secondDataSet *DataSet) (bool, error) {
	for i := 0; i < firstDataSet.getObjectCount(); i++ {
		differentObjectCount := 0
		object := firstDataSet.getObjects()[i]

		for j := 0; j < secondDataSet.getObjectCount(); j++ {
			objectCompared := secondDataSet.getObjects()[j]

			if isEqual, _ := object.Compare(objectCompared); isEqual == false {
				differentObjectCount++
			} else {
				firstDataSet.objects = remove(firstDataSet.getObjects(), i)
				secondDataSet.objects = remove(secondDataSet.getObjects(), j)
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
