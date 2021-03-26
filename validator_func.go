package sdk

type ValidatorFunc func(*Resource,*Operation,*Change) error
