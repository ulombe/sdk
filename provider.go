package sdk

type Provider interface {
    Name() string
    Provides() map[string]interface{}
    Execute(...string, ...Argument) (*Command, error)
}
