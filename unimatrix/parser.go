package unimatrix

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Parser struct {
	ReturnedName           string
	ReturnedTypeName       string
	Keys                   []string
	ReturnedResources      []Resource
	ReturnedCount          int
	ReturnedUnlimitedCount int
	ReturnedOffset         int
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

type ResourceIndex map[string]map[string]Resource

type AssociationIndex map[string]map[string]map[string][]string

type ResourceId struct {
	Id string `json:"id"`
}

var resourceIndex = make(ResourceIndex)
var associationIndex = make(AssociationIndex)
var resourceErrors []ResourceError

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

func buildResourceIndex(jsonResponse JsonResponse, associationIndex AssociationIndex, errors []ResourceError) {
	for responseKey, responseValue := range jsonResponse {
		if string([]rune(responseKey)[0]) != "$" && string(responseKey) != "errors" {
			var attributesListRaw []*json.RawMessage
			json.Unmarshal(*responseValue, &attributesListRaw)

			resourceIndex[responseKey] = make(map[string]Resource)

			for _, attributesRaw := range attributesListRaw {
				var resourceId ResourceId
				json.Unmarshal(*attributesRaw, &resourceId)

				resourceIndex[responseKey][resourceId.Id] = *NewResource(responseKey, attributesRaw)
			}
		}
	}
}

func buildAssociationIndex(staticResponse StaticResponse) {
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

func NewParser(rawResponse []byte) (*Parser, error) {
	var staticResponse StaticResponse
	var jsonResponse JsonResponse
	var associationIndex AssociationIndex
	var ids []string

	json.Unmarshal([]byte(rawResponse), &staticResponse)
	json.Unmarshal([]byte(rawResponse), &jsonResponse)

	if jsonResponse == nil {
		return nil, NewUnimatrixError("Unable to parse json response")
	}

	this := staticResponse.This
	ids = parseIds(this.Ids)
	resourceErrors = staticResponse.Errors

	buildAssociationIndex(staticResponse)
	buildResourceIndex(jsonResponse, associationIndex, staticResponse.Errors)

	return &Parser{
		ReturnedName:           this.Name,
		ReturnedTypeName:       this.TypeName,
		Keys:                   ids,
		ReturnedResources:      resources(this.Name, ids),
		ReturnedCount:          this.Count,
		ReturnedUnlimitedCount: this.UnlimitedCount,
		ReturnedOffset:         this.Offset,
	}, nil
}

func (parser *Parser) Name() (string, error) {
	return parser.ReturnedName, nil
}

func (parser *Parser) TypeName() (string, error) {
	return parser.ReturnedTypeName, nil
}

func (parser *Parser) Ids() ([]string, error) {
	return parser.Keys, nil
}

func (parser *Parser) Resources() ([]Resource, error) {
	return parser.ReturnedResources, nil
}

func (parser *Parser) Count() (int, error) {
	return parser.ReturnedCount, nil
}

func (parser *Parser) UnlimitedCount() (int, error) {
	return parser.ReturnedUnlimitedCount, nil
}

func (parser *Parser) Offset() (int, error) {
	return parser.ReturnedOffset, nil
}
