package sdk

import (
	"errors"
)

type State struct {
	name   string
	values AttributeValueMap
}

func NewState(name string) *State {
	return &State{
		name:   name,
		values: make(AttributeValueMap, 0),
	}
}

func (s *State) Name() string {
	return s.name
}

func (s *State) Data() AttributeValueMap {
	return s.values
}

func (s *State) Set(key string, value interface{}) {
	s.values.Set(key, value)
}

func (s *State) Merge(kv AttributeValueMap) {
	s.values.Merge(kv)
}

func (s *State) Keys() []string {
	return s.values.Keys()
}

func (s *State) Get(key string) (interface{}, error) {
	if state, ok := s.values[key]; ok {
		return state, nil
	}

	return nil, errors.New("Didn't find [key]")
}
