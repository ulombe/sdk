func RequiredInteger() *attributeType {
	return &attributeType{
		required: true,
		name:     "integer",
		value:    reflect.Int,
	}
}

func OptionalInteger() *attributeType {
	return &attributeType{
		required: false,
		name:     "integer",
		value:    reflect.Int,
	}
}
