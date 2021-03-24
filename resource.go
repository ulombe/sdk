package sdk

type Resource struct {
	name string
	provider *Provider
}

func NewResource(n string, p provider) *Resource {
    return &Resource{
        name: n,
				provider: p,
    }
}

func Provider() *Provider {
	return provider
}
