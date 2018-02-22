package unimatrix

type Resource map[string]string

func NewResource(name string, resource Resource) *Resource {
	return &resource
}
