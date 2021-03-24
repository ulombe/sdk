package provider

type ArgumentValidator func(*Argument) error

type Argument struct {
	Name string
	Value interface{}
	validators []ArgumentValidator
}

func (a *Argument) AddArgumentValidator(validators ...ArgumentValidator) {
	if a.validators == nil {
		a.validators := make([]ArgumentValidator, 0)
	}

	a.validators = append(a.validators, validators...)
}

func StringValidator(a *Argument) error {
	// return true if Value is string
}

func Int32Validator(a *Argument) error {

}

func Float64Validator(a *Argument) error {

}

func (a *Argument) Validate() error {
	for _, validator := range a.validators {
		if err := validator(a); err != nil {
			return err
		}
	}

	return nil
}
