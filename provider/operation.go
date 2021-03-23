package provider

type Operation interface {
	Render() ([]byte, error)
}
