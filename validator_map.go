package sdk

type ValidatorFunc func(*Resource,*Operation,*Change) error
type ValidatorMap map[*Operation]ValidatorFunc

func InitValidatorMap() ValidatorMap {
	return make(ValidatorMap, 0)
}
