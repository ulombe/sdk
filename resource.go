package sdk

type Resource struct {
	Name string
	Providers *ProviderList
}

func NewResource(name string) *Resource {
    return &Resource{
        Name: name,
    }
}
