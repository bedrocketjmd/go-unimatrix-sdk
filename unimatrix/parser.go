package unimatrix

import (
	"encoding/json"
)

type Parser struct {
	Name      string
	TypeName  string
	Keys      []string
	Resources []Resource
}

type JsonResponse map[string]*json.RawMessage

type StaticResponse struct {
	This struct {
		Name           string   `json:"name"`
		TypeName       string   `json:"type_name"`
		Ids            []string `json:"ids"`
		UnlimitedCount int      `json:"unlimited_count"`
		Offset         int      `json:"offset"`
		Count          int      `json:"count"`
	} `json:"$this"`
	AssociationTypes map[string][]*json.RawMessage `json:"$associations"`
}

type AssociationInner struct {
	TypeName string   `json:"type_name"`
	Ids      []string `json:"ids"`
}

type Associations map[string][]map[string]string

var associationMap = make(map[string]map[string]map[string][]string)
var resourceIndex = make(map[string]map[string]map[string]string)

func buildResourceIndex(jsonResponse JsonResponse) {
	for responseKey, responseValue := range jsonResponse {
		if string([]rune(responseKey)[0]) != "$" {
			var resources []map[string]string
			json.Unmarshal(*responseValue, &resources)

			resourceIndex[responseKey] = make(map[string]map[string]string)

			for _, resource := range resources {
				resourceIndex[responseKey][resource["id"]] = resource
			}
		}
	}
}

func buildAssociationMap(staticResponse StaticResponse) {
	associationTypes := staticResponse.AssociationTypes

	for associationType, _ := range associationTypes {
		associationMap[associationType] = make(map[string]map[string][]string)

		for _, associationList := range associationTypes[associationType] {
			var associationOuter map[string]*json.RawMessage
			json.Unmarshal(*associationList, &associationOuter)

			var associationOuterId string
			json.Unmarshal(*associationOuter["id"], &associationOuterId)

			associationMap[associationType][associationOuterId] = make(map[string][]string)

			for key, value := range associationOuter {
				if key != "id" {
					var associationInner map[string]*json.RawMessage
					json.Unmarshal(*value, &associationInner)

					var associationInnerIds []string
					json.Unmarshal(*associationInner["ids"], &associationInnerIds)

					associationMap[associationType][associationOuterId][key] = associationInnerIds
				}
			}
		}
	}
}

func (parser *Parser) GetAssociations(name string, id string) Associations {
	var associations = make(map[string][]map[string]string)

	for associationType, ids := range associationMap[name][id] {
		associations[associationType] = []map[string]string{}

		for _, id := range ids {
			associations[associationType] = append(associations[associationType], resourceIndex[associationType][id])
		}
	}

	return associations
}

func resources(name string, ids []string) []Resource {
	var resources []Resource

	for _, id := range ids {
		resources = append(resources, *NewResource(name, resourceIndex[name][id]))
	}

	return resources
}

func NewParser(rawResponse []byte) *Parser {
	var staticResponse StaticResponse
	var jsonResponse JsonResponse
	var ids []string

	json.Unmarshal([]byte(rawResponse), &staticResponse)
	json.Unmarshal([]byte(rawResponse), &jsonResponse)

	this := staticResponse.This
	name := this.Name
	typeName := this.TypeName
	ids = this.Ids

	buildResourceIndex(jsonResponse)
	buildAssociationMap(staticResponse)

	return &Parser{
		Name:      name,
		TypeName:  typeName,
		Keys:      ids,
		Resources: resources(name, ids),
	}
}
