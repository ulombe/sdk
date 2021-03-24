package provider

type Operation interface {
  Renderer
  Commander
}


type Renderer interface {
  Render() ([]byte, error)
}

type Commander interface {
  Command() string
}

type Validable interface {
  Validate() error
  AddValidator(...Validator)
}

type Validator func(*Validable) error
