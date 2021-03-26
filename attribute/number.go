package attribute

import (
  "gitlab.com/ulombe/sdk"
)

func OptionalNumber() *AttributeType {
	return sdk.NewAttributeType(false, "string", reflect.Float64)
}

func Number() *AttributeType {
	return sdk.NewAttributeType(true, "string", reflect.Float64)
}
