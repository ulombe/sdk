package attribute

import (
  "gitlab.com/ulombe/sdk"
)

func OptionalArray() *AttributeType {
	return sdk.NewAttributeType(false, "array", reflect.Slice)
}

func Array() *AttributeType {
	return sdk.NewAttributeType(true, "array", reflect.Slice)
}
