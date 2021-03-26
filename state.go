package sdk

// state.go
type State struct {
	name   string
	values Attributes
}

func NewState(name string) *State {
	return &State{
		name:   name,
		values: make(Attributes, 0),
	}
}

func (s *State) Set(key string, value interface{}) {
	s.values[key] = value
}
