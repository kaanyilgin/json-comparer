package main

import (
	"fmt"
	"os"

	"kaanyilgin.com/dataComparer/application"
	"kaanyilgin.com/dataComparer/infrastructure"
	"kaanyilgin.com/dataComparer/model"
)

func main() {
	firstFile := os.Args[1]
	secondFile := os.Args[2]

	reader := &infrastructure.FileReader{}
	dataSetComparer := &model.HashMapDataSetComparer{}
	jsonParser := &model.JSONObjectMapParser{}
	dataSetCompare := &application.DataSetCompare{
		DataReader:     reader,
		JSONParser:     jsonParser,
		DataSetFactory: model.InitDefaultDataSetFactory(dataSetComparer),
	}

	isEqaul, error := dataSetCompare.CompareDataSets(firstFile, secondFile, 1)

	if error != nil {
		fmt.Print(error)
		return
	}
	if isEqaul {
		fmt.Print("Files are equal")
	} else {
		fmt.Print("Files are not equal")
	}
}
