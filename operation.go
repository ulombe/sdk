package sdk

type Operation struct {
	name        string
	desiredState       string
	description string
}

func NewOperation(name string, desiredState string) *Operation {
	return &Operation{
		name:  name,
		desiredState: desiredState,
	}
}

func (o *Operation) Name() string {
	return o.name
}

func (o *Operation) DesiredState() string {
	return o.desiredState
}
