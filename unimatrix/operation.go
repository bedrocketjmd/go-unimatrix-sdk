package unimatrix

type Operation struct {
	path string
	// params map[string]string
}

func NewOperation(path string) *Operation {
	return &Operation{path: path}
}

func (operation *Operation) Read() ([]Resource, error) {
	response, error := Request(operation.path, "GET")

	if error != nil {
		return nil, error
	}

	return response.resources, nil
}
