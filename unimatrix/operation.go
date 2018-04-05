package unimatrix

type Operation struct {
	url        string
	parameters map[string][]string
}

func NewOperation(path string) *Operation {
	url := URL() + path
	return &Operation{url: url, parameters: map[string][]string{}}
}

func NewRealmScopedOperation(realm, resource string) *Operation {
	path := "/realms/" + realm + "/" + resource
	return NewOperation(path)
}

func (operation *Operation) Read() (*Response, error) {
	return Request(operation.url, "GET", operation.parameters, nil)
}

func (operation *Operation) Write(body interface{}) (*Response, error) {
	return Request(operation.url, "POST", operation.parameters, body)
}

func (operation *Operation) WriteResource(node string, resource Resource) (*Response, error) {
	var body = make(map[string][]interface{})
	var resources []interface{}
	resourceAttributes, _ := resource.Attributes()
	resources = append(resources, resourceAttributes)
	body[node] = resources

	return operation.Write(body)
}

func (operation *Operation) WriteResources(node string, resources []Resource) (*Response, error) {
	var body = make(map[string][]interface{})
	var bodyResources []interface{}
	for _, resource := range resources {
		resourceAttributes, _ := resource.Attributes()
		bodyResources = append(bodyResources, resourceAttributes)
	}
	body[node] = bodyResources

	return operation.Write(body)
}

func (operation *Operation) Destroy() (*Response, error) {
	return Request(operation.url, "DELETE", operation.parameters, nil)
}

func (operation *Operation) DestroyByUUID(uuid string) (*Response, error) {
	operation.AppendParameters(map[string][]string{"uuid": []string{uuid}})
	return operation.Destroy()
}

func (operation *Operation) DestroyByUUIDs(uuids []string) (*Response, error) {
	operation.AppendParameters(map[string][]string{"uuid:in[]": uuids})
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

func (operation *Operation) SetAccessToken(token string) {
	operation.AppendParameters(map[string][]string{"access_token": []string{token}})
}
