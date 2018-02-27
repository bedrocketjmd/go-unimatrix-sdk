package unimatrix

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Parser struct {
	Name      string
	TypeName  string
	Keys      []string
	Resources []Resource
	Errors    []map[string]interface{}
}

type JsonResponse map[string]*json.RawMessage

type StaticResponse struct {
	This struct {
		Name           string        `json:"name"`
		TypeName       string        `json:"type_name"`
		Ids            []interface{} `json:"ids"`
		UnlimitedCount int           `json:"unlimited_count"`
		Offset         int           `json:"offset"`
		Count          int           `json:"count"`
	} `json:"$this"`
	AssociationTypes map[string][]*json.RawMessage `json:"$associations"`
}

type AssociationInner struct {
	TypeName string   `json:"type_name"`
	Ids      []string `json:"ids"`
}

type Associations map[string][]map[string]interface{}

var associationMap = make(map[string]map[string]map[string][]string)
var resourceIndex = make(map[string]map[string]map[string]interface{})

func parseIds(idsInterface []interface{}) []string {
	var ids []string

	for _, idInterface := range idsInterface {
		switch id := idInterface.(type) {
		case float64:
			ids = append(ids, fmt.Sprintf("%.0f", id))
		case int:
			ids = append(ids, strconv.Itoa(id))
		case string:
			ids = append(ids, id)
		}
	}

	return ids
}

func parseResource(resourceJson map[string]*json.RawMessage) map[string]string {
	var resource = make(map[string]string)

	for key, value := range resourceJson {
		var attribute string
		var attributeInterface interface{}

		json.Unmarshal(*value, &attributeInterface)

		switch typedValue := attributeInterface.(type) {
		case float64:
			attribute = fmt.Sprintf("%.0f", typedValue)
		case int:
			attribute = strconv.Itoa(typedValue)
		default:
			json.Unmarshal(*value, &attribute)
		}

		resource[key] = attribute
	}

	return resource
}

func buildResourceIndex(jsonResponse JsonResponse) {
	for responseKey, responseValue := range jsonResponse {
		if string([]rune(responseKey)[0]) != "$" {
			var resources []map[string]interface{}

			json.Unmarshal(*responseValue, &resources)

			resourceIndex[responseKey] = make(map[string]map[string]interface{})

			for _, resource := range resources {
				resourceIndex[responseKey][resource["id"].(string)] = resource
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
			var associationOuterId string

			json.Unmarshal(*associationList, &associationOuter)
			json.Unmarshal(*associationOuter["id"], &associationOuterId)

			associationMap[associationType][associationOuterId] = make(map[string][]string)

			for key, value := range associationOuter {
				if key != "id" {
					var associationInner map[string]*json.RawMessage
					var associationInnerIds []string

					json.Unmarshal(*value, &associationInner)
					json.Unmarshal(*associationInner["ids"], &associationInnerIds)

					associationMap[associationType][associationOuterId][key] = associationInnerIds
				}
			}
		}
	}
}

func resources(name string, ids []string) []Resource {
	var resources []Resource

	for _, id := range ids {
		resources = append(resources, *NewResource(name, resourceIndex[name][id]))
	}

	return resources
}

func errors(staticResponse StaticResponse) []map[string]interface{} {
	var errors []map[string]interface{}

	for _, resource := range resourceIndex["errors"] {
		errors = append(errors, resource)
	}

	return errors
}

func (parser *Parser) GetAssociations(name string, id string) Associations {
	var associations = make(map[string][]map[string]interface{})

	for associationType, ids := range associationMap[name][id] {
		for _, id := range ids {
			associations[associationType] = append(associations[associationType], resourceIndex[associationType][id])
		}
	}

	return associations
}

func NewParser(rawResponse []byte) *Parser {
	var staticResponse StaticResponse
	var jsonResponse JsonResponse
	var ids []string

	json.Unmarshal([]byte(rawResponse), &staticResponse)
	json.Unmarshal([]byte(rawResponse), &jsonResponse)

	this := staticResponse.This
	ids = parseIds(this.Ids)

	buildResourceIndex(jsonResponse)
	buildAssociationMap(staticResponse)

	return &Parser{
		Name:      this.Name,
		TypeName:  this.TypeName,
		Keys:      ids,
		Resources: resources(this.Name, ids),
		Errors:    errors(staticResponse),
	}
}
