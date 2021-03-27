package sdk

import (
	"reflect"
	"testing"
)

func TestChange_String(t *testing.T) {
	tt := []struct{
		name string
		got AttributeValueMap
		expected map[string]interface{}
	}{
		{
			name: "Create User#1",
			got: AttributeValueMap {
				"name": "John",
				"surname": "Doe",
			},
			expected: map[string]interface{} {
				"name": "John",
				"surname": "Doe",
			},
		},
		{
			name: "Create User#2",
			got: AttributeValueMap {
				"name": "John",
				"age": 29,
				"tall": true,
				"fullName": map[string]string{
					"first": "John",
					"last": "Doe",
				},
			},
			expected: map[string]interface{} {
				"name": "John",
				"age": 29,
				"tall": true,
				"fullName": map[string]string{
					"first": "John",
					"last": "Doe",
				},
			},
		},
		{
			name: "Create User#3",
			got: AttributeValueMap {
				"name": "John",
				"age": 29,
				"tall": nil,
			},
			expected: map[string]interface{} {
				"name": "John",
				"age": 29,
				"tall": nil,
			},
		},

	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			newChange := NewChange(test.got)

			for k, _ := range test.got {
				if attr, err := newChange.Get(k); err != nil {
					t.Errorf("Didn't find a value for %s key", k)
				} else {
					if !reflect.DeepEqual(attr, test.expected[k]) {
						t.Errorf("Expected %v for attribute %v but found %v", test.expected[k], k, attr)
					}
				}
			}
		})
	}
}