package provider

type Operation interface {
	Script() ([]byte, error)
}
