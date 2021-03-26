package sdk

type Change struct {
	values Attributes
}

func (c *Change) Get(key string) (attributeValue, error) {
	if value, ok := c.values[key]; ok {
		return value, nil
	}

	return nil, errors.New(fmt.Sprintf("%s doesn't exist", key))
}

func NewChange(attributes Attributes) *Change {
	newChange := new(Change)
	newChange.values = attributes

	return newChange
}
