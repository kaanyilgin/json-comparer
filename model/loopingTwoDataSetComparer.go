package model

//LoopingTwoDataSetComparer compares two dataset going through whole dataset
type LoopingTwoDataSetComparer struct {
	ObjectComparer ObjectComparer
}

//Compare compares two dataset
func (l LoopingTwoDataSetComparer) Compare(firstDataSet *DataSet, secondDataSet *DataSet) bool {
	firstDataSetObjects := firstDataSet.objects.(*JSONObjectArray)
	secondDataSetObjects := secondDataSet.objects.(*JSONObjectArray)

	for i := 0; i < firstDataSetObjects.GetLength(); i++ {
		differentObjectCount := 0
		object := firstDataSetObjects.dictonary[i]

		for j := 0; j < secondDataSetObjects.GetLength(); j++ {
			objectCompared := secondDataSetObjects.dictonary[j]

			if l.ObjectComparer.Compare(object, objectCompared) == false {
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
				return false
			}
		}
	}

	return true
}

func remove(slice []JSONObject, s int) []JSONObject {
	return append(slice[:s], slice[s+1:]...)
}
