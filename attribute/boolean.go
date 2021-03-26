package attribute

import (
  "gitlab.com/ulombe/sdk"
)

func OptionalBoolean() *AttributeType {
	return sdk.NewAttributeType(false, "boolean", reflect.Bool)
}

func Boolean() *AttributeType {
	return sdk.NewAttributeType(true, "boolean", reflect.Bool)
}
