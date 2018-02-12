package unimatrix_sdk

// var typeRegistry = make(map[string]reflect.Type)
//
// func registerType(myType interface{}) {
// 	t := reflect.TypeOf(myType).Elem()
// 	typeRegistry[t.Name()] = t
// }
//
// type Artifact struct {
// 	name string
// }
//
// func init() {
// 	registerType((*Artifact)(nil))
// }
//
// func makeInstance(name string) interface{} {
// 	v := reflect.New(typeRegistry[name]).Elem()
// 	// Maybe fill in fields here if necessary
// 	return v.Interface()
// }

type Resource struct {
	attributes map[string]interface{}
}

func NewResource(name string, r interface{}) *Resource {
	return &Resource{attributes: r.(map[string]interface{})}
}
