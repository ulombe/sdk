package attribute

import (
  "gitlab.com/ulombe/sdk"
)

func OptionalObject() *AttributeType {
	return sdk.NewAttributeType(false, "string", reflect.Map)
}

func Object() *AttributeType {
	return sdk.NewAttributeType(true, "string", reflect.Map)
}
