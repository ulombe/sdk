package sdk

import (
	"reflect"
	"testing"
)

func TestAttributeType(t *testing.T) {
	type data struct {
		name string
		required bool
		kind reflect.Kind
	}

	type expected struct {
		name string
		required bool
		kind reflect.Kind
	}

	tt := []struct{
		desc string
		data data
		expected expected
	}{
		{
			desc: "required integer",
			data: data{
				name: "integer",
				required: true,
				kind: reflect.Int,
			},
			expected: expected {
				name: "integer",
				required: true,
				kind: reflect.Int,
			},
		},
		{
			desc: "optional integer",
			data: data{
				name: "integer",
				required: false,
				kind: reflect.Int,
			},
			expected: expected {
				name: "integer",
				required: false,
				kind: reflect.Int,
			},
		},
		{
			desc: "required number",
			data: data{
				name: "number",
				required: true,
				kind: reflect.Float64,
			},
			expected: expected {
				name: "number",
				required: true,
				kind: reflect.Float64,
			},
		},
		{
			desc: "optional number",
			data: data{
				name: "number",
				required: false,
				kind: reflect.Float64,
			},
			expected: expected {
				name: "number",
				required: false,
				kind: reflect.Float64,
			},
		},
		{
			desc: "required string",
			data: data{
				name: "string",
				required: true,
				kind: reflect.String,
			},
			expected: expected {
				name: "string",
				required: true,
				kind: reflect.String,
			},
		},
		{
			desc: "optional string",
			data: data{
				name: "string",
				required: false,
				kind: reflect.String,
			},
			expected: expected {
				name: "string",
				required: false,
				kind: reflect.String,
			},
		},
		{
			desc: "required boolean",
			data: data{
				name: "boolean",
				required: true,
				kind: reflect.Bool,
			},
			expected: expected {
				name: "boolean",
				required: true,
				kind: reflect.Bool,
			},
		},
		{
			desc: "optional boolean",
			data: data{
				name: "boolean",
				required: false,
				kind: reflect.Bool,
			},
			expected: expected {
				name: "boolean",
				required: false,
				kind: reflect.Bool,
			},
		},
		{
			desc: "required object",
			data: data{
				name: "object",
				required: true,
				kind: reflect.Map,
			},
			expected: expected {
				name: "object",
				required: true,
				kind: reflect.Map,
			},
		},
		{
			desc: "optional object",
			data: data{
				name: "object",
				required: false,
				kind: reflect.Map,
			},
			expected: expected {
				name: "object",
				required: false,
				kind: reflect.Map,
			},
		},

		{
			desc: "required array",
			data: data{
				name: "array",
				required: true,
				kind: reflect.Slice,
			},
			expected: expected {
				name: "array",
				required: true,
				kind: reflect.Slice,
			},
		},
		{
			desc: "optional array",
			data: data{
				name: "array",
				required: false,
				kind: reflect.Slice,
			},
			expected: expected {
				name: "array",
				required: false,
				kind: reflect.Slice,
			},
		},

	}

	for _, test := range tt {
		t.Run(test.desc, func(t *testing.T) {
			attrType := NewAttributeType(test.data.required, test.data.name, test.data.kind)

			if attrType.Name() != test.expected.name {
				t.Errorf("attrType name expected %v got %v", test.expected.name, attrType.Name())
			}
			
			if attrType.Required() != test.expected.required {
				t.Errorf("attrType name expected %v got %v", test.expected.required, attrType.Required())
			}

			if attrType.Kind() != test.expected.kind {
				t.Errorf("attrType name expected %v got %v", test.expected.kind, attrType.Kind())
			}
			
		})
	}
}