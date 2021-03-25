package sdk

import (
  "gitlab.com/ulombe/sdk/provider"
)

type Provider func(...string, ...provider.Argument) (provider.Operation, error)
