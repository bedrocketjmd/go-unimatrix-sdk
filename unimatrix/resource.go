package unimatrix

type Resource map[string]interface{}

func NewResource(name string, resource Resource) *Resource {
	return &resource
}
