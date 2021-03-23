package sdk

type ProviderList struct {
	providers map[string]*Provider
}

func(p *ProviderList) Add(newProvider *Provider) {
	p.providers[newProvider.Name()] = newProvider
}

func NewProviderList(providers ...Provider) *ProviderList {
	list := new(ProviderList)
	list.providers = make([map[string]*Provider, 0)

	for _, provider := range providers {
		list.Add(provider)
	}

	return list
}
