package sdk

import (
	"sort"
)

type AttributeValueMap map[string]attributeValue

func NewAttributeValueMap() AttributeValueMap {
	return map[string]attributeValue{}
}

func (a AttributeValueMap) Set(k string, value interface{}) {
	a[k] = value
}

func (a AttributeValueMap) Keys() []string {
	keys := make([]string, 0)

	for k, _ := range a {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (a AttributeValueMap) Merge(kv AttributeValueMap) {
	for k, v := range kv {
		a[k] = v
	}
}