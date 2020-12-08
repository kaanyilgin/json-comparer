package model

import "strings"

// ParseJSON parse json into a dynamic object
func ParseJSON(data string) []string {
	data = strings.ReplaceAll(data, " ", "")
	data = strings.Replace(data, "[", "", 1)
	data = strings.Replace(data, "]", "", 1)
	splittedObject := strings.Split(data, ",")
	return splittedObject
}
