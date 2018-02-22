package unimatrix

type Resource struct {
	attributes map[string]interface{}
}

func NewResource(name string, resource interface{}) *Resource {
	return &Resource{attributes: resource.(map[string]interface{})}
}
