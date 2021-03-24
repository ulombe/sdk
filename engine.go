package sdk

type ValidationError []error

type Engine struct {
  resources map[string]*Resource
}

func NewEngine() *Engine {
  engine := &Engine{}
  engine.resources = make(map[string]*Resource)

  return engine
}

func (e *Engine) PushResource(resource *Resource) {
  e.resources[r.Name()] = resource
}

func (e *Engine) Validate() error {
  errorList := make(error, 0)

  for name, resource := range resources {
    if err := resource.Provider().Validate(); err != nil {
      errorList = append(errorList, err)
    }
  }

  if len(errorList) == 0 {
    return nil
  }

  return ValidationError(errorList)
}
