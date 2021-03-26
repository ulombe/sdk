package sdk

type operator struct {
	resource *Resource
}

func NewOperator(r *Resource) *operator {
	return &operator{
		resource: r,
	}
}

func (o *operator) Apply(operation *Operation, changes ...*Change) ([]*State, error) {
	stateSet := make([]*State, 0)

	for _, c := range changes {

		state := NewState(operation.State())
		if rules, ok := o.resource.Attributes[operation]; ok {
			for key, rule := range rules {
				if rule.Required() {
					if value, err := c.Get(key); err == nil {
						if rule.Kind() == reflect.ValueOf(value).Kind() {
							state.Set(key, value)
						} else {
							return nil, errors.New(fmt.Sprintf("Kinds %s don't match %s", rule.Kind(), reflect.ValueOf(value).Kind()))
						}
					} else {
						return nil, errors.New(fmt.Sprintf("%s is required", key))
					}
				} else {
					if value, err := c.Get(key); err == nil {
						if rule.Kind() == reflect.ValueOf(value).Kind() {
							state.Set(key, value)
						} else {
							return nil, errors.New(fmt.Sprintf("Kinds %s don't match %s", rule.Kind(), reflect.ValueOf(value).Kind()))
						}
					}
				}
			}

      stateSet = append(stateSet, state)
		}
	}

	return stateSet, nil
}
