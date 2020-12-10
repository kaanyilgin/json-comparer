package infrastructure

import (
	"io/ioutil"
)

// DataReader provides an interface to read data from sources
type DataReader interface {
	Read(fileName string) (string, error)
}

//FileReader reads file
type FileReader struct {
	DataReader
}

//Read returns the file content
func (f FileReader) Read(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}
