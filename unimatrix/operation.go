package unimatrix

type Operation struct {
	path       string
	parameters map[string]string
}

func NewOperation(path string) *Operation {
	return &Operation{path: path}
}

func (operation *Operation) Read() ([]Resource, error) {
	response, error := Request(operation.path, "GET", operation.parameters)

	if error != nil {
		return nil, error
	}

	return response.resources, nil
}

func (operation *Operation) SetParameters(parameters map[string]string) {
	operation.parameters = parameters
}
