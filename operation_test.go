package sdk

import (
	"testing"
)

func TestOperation(t *testing.T) {
	type data struct {
		name string
		desiredState string
	}

	type expected struct {
		name string 
		desiredState string
	}

	type testSpec struct {
		desc string
		data data
		expected expected
	}

	tt := []testSpec{
		{
			desc: "Create Operation",
			data: data{
				name: "create",
				desiredState: "present",
			},
			expected: expected{
				name: "create",
				desiredState: "present",
			},
		},
		{
			desc: "Delete Operation",
			data: data{
				name: "delete",
				desiredState: "absent",
			},
			expected: expected{
				name: "delete",
				desiredState: "absent",
			},
		},
		{
			desc: "Update Operation",
			data: data{
				name: "update",
				desiredState: "present",
			},
			expected: expected{
				name: "update",
				desiredState: "present",
			},
		},

	}

	for _, test := range tt {
		t.Run(test.desc, func(t *testing.T) {
			operation := NewOperation(test.data.name, test.data.desiredState)
			if operation.Name() != test.expected.name {
				t.Errorf("Expected %v but got %v for", test.expected.name, operation.Name())
			}

			if operation.DesiredState() != test.expected.desiredState {
				t.Errorf("Expected %v but got %v for", test.expected.desiredState, operation.DesiredState())
			}
		})
	}
}