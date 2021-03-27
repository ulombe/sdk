package sdk

import (
	"reflect"
	"testing"
)

func TestAttributeValueMap_Set(t *testing.T) {
	type data struct {
		value map[string]string
	}

	tt := []struct{
		name string
		data data
		expected []string
	}{
		{
			name: "person",
			data: data {
				value: map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",	
				},
			},
			expected: []string{"key1", "key2", "key3"},
		},
		{
			name: "person",
			data: data {
				value: map[string]string{
					"a": "value1",
					"b": "value2",
					"c": "value3",
					"d": "value3",
					"e": "value3",
					"f": "value3",	
				},
			},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			attributeMap := NewAttributeValueMap()
			for k, v := range test.data.value {
				attributeMap.Set(k, v)
			}

			if !reflect.DeepEqual(attributeMap.Keys(), test.expected) {
				t.Errorf("newSate shoud have %q keys but got %q", attributeMap.Keys(), test.expected)
			}
		})
	}
}

func TestAttributeValueMap_Merge(t *testing.T) {
	type data struct {
		value AttributeValueMap
	}

	tt := []struct{
		name string
		data data
		expected []string
	}{
		{
			name: "person",
			data: data {
				value: AttributeValueMap{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",	
				},
			},
			expected: []string{"key1", "key2", "key3"},
		},
		{
			name: "person",
			data: data {
				value: AttributeValueMap{
					"a": "value1",
					"b": "value2",
					"c": "value3",
					"d": "value3",
					"e": "value3",
					"f": "value3",	
				},
			},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			attributeMap := NewAttributeValueMap()
			attributeMap.Merge(test.data.value)

			if !reflect.DeepEqual(attributeMap.Keys(), test.expected) {
				t.Errorf("attributeMap shoud have %q keys but got %q", attributeMap.Keys(), test.expected)
			}
		})
	}
}