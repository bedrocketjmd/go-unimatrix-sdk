package unimatrix

type Operation struct {
	path       string
	parameters string
}

func NewOperation(path string) *Operation {
	return &Operation{path: path}
}

func (operation *Operation) Read() (*Parser, error) {
	URL := GetURL() + operation.path

	response, error := Request(URL, "GET", operation.parameters)

	if error != nil {
		return nil, error
	}

	return response, nil
}

func (operation *Operation) SetParameters(parameters string) {
	operation.parameters = parameters
}
