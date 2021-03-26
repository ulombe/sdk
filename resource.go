package sdk

// resource.go
type Resource struct {
	Attributes map[*Operation]AttributeTypeMap
	Validators map[*Operation]func(*Resource,*Operation,*Change) error
}

func (r *Resource) Operations() []string {
	operations := make([]string, 0)
	for k, _ := range r.Attributes {
		operations = append(operations, k.Name())
	}

	return operations

}

func (r *Resource) Validator(o *Operation, validate func(res *Resource, o *Operation, c *Change) error) {
	r.Validators[o] = validate
}

func (r *Resource) Validate(o *Operation, c *Change) error {
	return r.Validators[o](r, o, c)
}
