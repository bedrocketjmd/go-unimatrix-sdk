package unimatrix

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Parser struct {
	Name           string
	TypeName       string
	Keys           []string
	Resources      []Resource
	Count          int
	UnlimitedCount int
	Offset         int
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

func buildResourceIndex(jsonResponse JsonResponse, associationIndex AssociationIndex, errors []ResourceError) ResourceIndex {
	var resourceIndex = make(ResourceIndex)

	for responseKey, responseValue := range jsonResponse {
		if string([]rune(responseKey)[0]) != "$" && string(responseKey) != "errors" {
			var attributesListRaw []*json.RawMessage
			json.Unmarshal(*responseValue, &attributesListRaw)

			resourceIndex[responseKey] = make(map[string]Resource)

			for _, attributesRaw := range attributesListRaw {
				var resourceId ResourceId
				json.Unmarshal(*attributesRaw, &resourceId)

				resource := *NewResource(responseKey, attributesRaw)
				resource.AddAssociationIndices(&resourceIndex, &associationIndex, &errors)
				resourceIndex[responseKey][resourceId.Id] = resource

			}
		}
	}

	return resourceIndex
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

func resources(name string, ids []string, resourceIndex ResourceIndex) []Resource {
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
	var resourceIndex ResourceIndex
	var resourceErrors []ResourceError
	var ids []string

	json.Unmarshal([]byte(rawResponse), &staticResponse)
	json.Unmarshal([]byte(rawResponse), &jsonResponse)

	if jsonResponse == nil {
		return nil, NewUnimatrixError("Unable to parse json response")
	}

	this := staticResponse.This
	ids = parseIds(this.Ids)
	resourceErrors = staticResponse.Errors
	associationIndex = buildAssociationIndex(staticResponse)
	resourceIndex = buildResourceIndex(jsonResponse, associationIndex, resourceErrors)

	return &Parser{
		Name:           this.Name,
		TypeName:       this.TypeName,
		Keys:           ids,
		Resources:      resources(this.Name, ids, resourceIndex),
		Count:          this.Count,
		UnlimitedCount: this.UnlimitedCount,
		Offset:         this.Offset,
	}, nil
}
