package unimatrix

import (
	"net/url"
)

type Operation struct {
	path       string
	parameters map[string][]string
}

func NewOperation(path string) *Operation {
	return &Operation{path: path, parameters: map[string][]string{}}
}

func NewRealmScopedOperation(realm, resource string) *Operation {
	path := "/realms/" + realm + "/" + resource
	return NewOperation(path)
}

func (operation *Operation) Read() (*Parser, error) {
	URL := GetURL() + operation.path

	var parameters url.Values
	parameters = operation.parameters

	response, error := Request(URL, "GET", parameters.Encode())

	if error != nil {
		return nil, error
	}

	return response, nil
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
