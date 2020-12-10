package model

import (
	"sort"
)

// AttributeSorter sort json attrbutes
type AttributeSorter interface {
	sort(attributes map[string]interface{}) map[string]interface{}
}

// AttributeAlphabeticalSorter sort attrbute sorter alphabetical
type AttributeAlphabeticalSorter struct {
}

// Sort compares the given json object
func (sorter AttributeAlphabeticalSorter) Sort(attributes map[string]interface{}) map[string]interface{} {
	attributesCount := len(attributes)
	keys := make([]string, attributesCount)

	i := 0
	for key := range attributes {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	sortedAttributes := make(map[string]interface{}, attributesCount)

	for i := range keys {
		sortedAttributes[keys[i]] = attributes[keys[i]]
	}

	return sortedAttributes
}
