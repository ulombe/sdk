package attribute

import (
  "gitlab.com/ulombe/sdk"
)

func OptionalString() *AttributeType {
	return sdk.NewAttributeType(false, "string", reflect.String)
}

func String() *AttributeType {
	return sdk.NewAttributeType(true, "string", reflect.String)
}
