package sdk

// operation.go
type Operation struct {
	name        string
	state       string
	description string
}

func NewOperation(name string, state string) *Operation {
	return &Operation{
		name:  name,
		state: state,
	}
}

func (o *Operation) Name() string {
	return o.name
}

func (o *Operation) State() strig {
	return o.state
}
