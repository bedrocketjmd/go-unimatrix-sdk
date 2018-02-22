package unimatrix

import (
	"encoding/json"
)

type Parser struct {
	name         string
	typeName     string
	keys         []string
	associations []Resource
	resources    []Resource
}

func NewParser(rawResponse []byte) *Parser {
	var response map[string]interface{}
	var resources []Resource
	var ids []string

	json.Unmarshal([]byte(rawResponse), &response)

	this := response["$this"].(map[string]interface{})
	name := this["name"].(string)
	typeName := this["type_name"].(string)

	for _, id := range this["ids"].([]interface{}) {
		ids = append(ids, id.(string))
	}

	for _, resource := range response[name].([]interface{}) {
		resources = append(resources, *NewResource(name, resource))
	}
	return &Parser{
		name:      name,
		typeName:  typeName,
		keys:      ids,
		resources: resources,
	}
}
