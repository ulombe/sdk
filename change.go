package sdk

import (
	"fmt"
	"errors"
)

type Change struct {
	values AttributeValueMap
}

func (c *Change) Get(key string) (interface{}, error) {
	if value, ok := c.values[key]; ok {
		return value, nil
	}

	return nil, errors.New(fmt.Sprintf("%s doesn't exist", key))
}

func NewChange(attributes AttributeValueMap) *Change {
	newChange := new(Change)
	newChange.values = attributes

	return newChange
}
