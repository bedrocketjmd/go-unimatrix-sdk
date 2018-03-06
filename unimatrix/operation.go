package unimatrix

type Operation struct {
	url        string
	parameters map[string][]string
}

func NewOperation(path string) *Operation {
	url := GetURL() + path
	return &Operation{url: url, parameters: map[string][]string{}}
}

func NewRealmScopedOperation(realm, resource string) *Operation {
	path := "/realms/" + realm + "/" + resource
	return NewOperation(path)
}

func (operation *Operation) Read() ([]Resource, error) {
	response, error := Request(operation.url, "GET", operation.parameters, nil)
	return response, error
}

func (operation *Operation) Write(node string, objects interface{}) ([]Resource, error) {
	var body = make(map[string]interface{})
	body[node] = objects

	response, error := Request(operation.url, "POST", operation.parameters, body)
	return response, error
}

func (operation *Operation) Destroy() ([]Resource, error) {
	response, error := Request(operation.url, "DELETE", operation.parameters, nil)
	return response, error
}

func (operation *Operation) DestroyByUUID(uuid string) ([]Resource, error) {
	operation.AppendParameters(map[string][]string{"uuid": []string{uuid}})
	return operation.Destroy()
}

func (operation *Operation) DestroyByUUIDs(uuids []string) ([]Resource, error) {
	operation.AppendParameters(map[string][]string{"uuid": uuids})
	return operation.Destroy()
}

func (operation *Operation) AssignParameters(parameters map[string][]string) {
	operation.parameters = parameters
}

func (operation *Operation) AppendParameters(parameters map[string][]string) {
	for parameter, values := range parameters {
		if parameter[len(parameter)-2:] == "[]" {
			newValues := append(operation.parameters[parameter], values...)
			operation.parameters[parameter] = newValues
		} else {
			operation.parameters[parameter] = values
		}
	}
}
