package sdk

type Provider func(...string, ...provider.Argument) (provider.Operation, error)
