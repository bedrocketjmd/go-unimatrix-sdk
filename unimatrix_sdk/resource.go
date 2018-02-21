package unimatrix_sdk

type Resource struct {
	attributes map[string]interface{}
}

func NewResource(name string, r interface{}) *Resource {
	return &Resource{attributes: r.(map[string]interface{})}
}
