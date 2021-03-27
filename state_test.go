package sdk

import (
  "reflect"
  "testing"
)

func TestState_Name(t *testing.T) {
  type data struct {
    name string
  }

  tt := []struct{
    name string
    values data
    expected string
  }{
    {
      name: "present",
      values: data {
        name: "present",
      },
      expected: "present",
    },
    {
      name: "absent",
      values: data {
        name: "absent",
      },
      expected: "absent",
    },
    {
      name: "unknown",
      values: data {
        name: "unknown",
      },
      expected: "unknown",
    },
  }

  for _, test := range tt {
    t.Run(test.name, func(t *testing.T) {
      newState := NewState(test.values.name)

      if newState.Name() != test.expected {
        t.Errorf("The state name should be %s but got %s", test.expected, newState.Name())
      }
    })
  }
}

func TestState_Set_Values(t *testing.T) {
  type data struct {
    name string
    values AttributeValueMap
  }

  tt := []struct{
    name string
    data data
    expected []string
  }{
    {
      name: "present",
      data: data{
        name: "present",
        values: AttributeValueMap{
          "key1": "value1",
          "key2": "value2",
          "key3": "value3",
        },
      },
      expected: []string{"key1", "key2", "key3"},
    },
    {
      name: "present",
      data: data{
        name: "present",
        values: AttributeValueMap{
          "key1": "value1",
          "key2": "value2",
        },
      },
      expected: []string{"key1", "key2"},
    },
  }

  for _, test := range tt {
    t.Run(test.name, func(t *testing.T) {
      newState := NewState(test.data.name)
      for k, v := range test.data.values {
        newState.Set(k, v)
      }

      if !reflect.DeepEqual(newState.Keys(), test.expected) {
        t.Errorf("newState shoud have %q keys but got %q", newState.Keys(), test.expected)
      }
    })
  }
}

func TestState_Merge_Values(t *testing.T) {
  type data struct {
    name string
    values AttributeValueMap
  }

  tt := []struct{
    name string
    data data
    expected []string
  }{
    {
      name: "present",
      data: data{
        name: "present",
        values: AttributeValueMap{
          "key1": "value1",
          "key2": "value2",
          "key3": "value3",
        },
      },
      expected: []string{"key1", "key2", "key3"},
    },
    {
      name: "present",
      data: data{
        name: "present",
        values: AttributeValueMap{
          "key1": "value1",
          "key2": "value2",
        },
      },
      expected: []string{"key1", "key2"},
    },
  }

  for _, test := range tt {
    t.Run(test.name, func(t *testing.T) {
      newState := NewState(test.data.name)
      newState.Merge(test.data.values)
      
      if !reflect.DeepEqual(newState.Keys(), test.expected) {
        t.Errorf("newState shoud have %q keys but got %q", newState.Keys(), test.expected)
      }
    })
  }
}
