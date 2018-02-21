package unimatrix

type Operation struct {
	path string
	// params map[string]string
}

func NewOperation(path string) *Operation {
	return &Operation{path: path}
}

func (o *Operation) Read() ([]Resource, error) {
	response, err := Request(o.path, "GET")

	if err != nil {
		return nil, err
	}

	return response.resources, nil
}
