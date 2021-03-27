package sdk

import (
	"reflect"
)

type AttributeType struct {
	required bool
	name     string
	value    reflect.Kind
}

func NewAttributeType(required bool, name string, value reflect.Kind) *AttributeType {
  return &AttributeType{
    required: required,
    name: name,
    value: value,
  }
}

func (a AttributeType) Name() string {
	return a.name
}

func (a AttributeType) Required() bool {
	return a.required
}

func (a AttributeType) Kind() reflect.Kind {
	return a.value
}
