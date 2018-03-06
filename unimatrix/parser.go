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
	Errors    []ResourceError
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
	Errors           []ResourceError               `json:"errors"`
}

type AssociationIndex map[string]map[string]map[string][]string

var resourceIndex = make(map[string]map[string]Resource)

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

func buildResourceIndex(name string, jsonResponse JsonResponse, associationIndex AssociationIndex, errors []ResourceError) {
	for responseKey, responseValue := range jsonResponse {
		if string([]rune(responseKey)[0]) != "$" && string(responseKey) != "errors" {
			var resourceAttributes []ResourceAttributes

			json.Unmarshal(*responseValue, &resourceAttributes)

			resourceIndex[responseKey] = make(map[string]Resource)

			for _, resourceAttributes := range resourceAttributes {
				var id = resourceAttributes["id"].(string)
				var resourceErrors []ResourceError

				if len(associationIndex[name][id]["errors"]) > 0 {
					for _, error := range errors {
						for _, errorId := range associationIndex[name][id]["errors"] {
							if error.Id == errorId {
								resourceErrors = append(resourceErrors, error)
								break
							}
						}
					}
				}

				resourceIndex[responseKey][id] = *NewResource(responseKey, resourceAttributes, associationIndex[name][id], resourceErrors)
			}
		}
	}
}

func buildAssociationIndex(staticResponse StaticResponse) AssociationIndex {
	var associationIndex = make(AssociationIndex)
	associationTypes := staticResponse.AssociationTypes

	for associationType, _ := range associationTypes {
		associationIndex[associationType] = make(map[string]map[string][]string)

		for _, associationList := range associationTypes[associationType] {
			var associationOuter map[string]*json.RawMessage
			var associationOuterId string

			json.Unmarshal(*associationList, &associationOuter)
			json.Unmarshal(*associationOuter["id"], &associationOuterId)

			associationIndex[associationType][associationOuterId] = make(map[string][]string)

			for key, value := range associationOuter {
				if key != "id" {
					var associationInner map[string]*json.RawMessage
					var associationInnerIds []string

					json.Unmarshal(*value, &associationInner)
					json.Unmarshal(*associationInner["ids"], &associationInnerIds)

					associationIndex[associationType][associationOuterId][key] = associationInnerIds
				}
			}
		}
	}

	return associationIndex
}

func resources(name string, ids []string) []Resource {
	var resources []Resource

	for _, id := range ids {
		if len(resourceIndex[name]) > 0 {
			resources = append(resources, resourceIndex[name][id])
		}
	}

	return resources
}

func NewParser(rawResponse []byte) *Parser {
	var staticResponse StaticResponse
	var jsonResponse JsonResponse
	var associationIndex AssociationIndex
	var ids []string

	json.Unmarshal([]byte(rawResponse), &staticResponse)
	json.Unmarshal([]byte(rawResponse), &jsonResponse)

	this := staticResponse.This
	ids = parseIds(this.Ids)
	associationIndex = buildAssociationIndex(staticResponse)

	buildResourceIndex(this.Name, jsonResponse, associationIndex, staticResponse.Errors)

	return &Parser{
		Name:      this.Name,
		TypeName:  this.TypeName,
		Keys:      ids,
		Resources: resources(this.Name, ids),
		Errors:    staticResponse.Errors,
	}
}
