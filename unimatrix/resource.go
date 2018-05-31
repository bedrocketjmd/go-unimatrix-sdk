package unimatrix

import "encoding/json"

type ResourceAssociations map[string][]Resource

type ResourceError struct {
	Id        string `json:"id"`
	TypeName  string `json:"type_name"`
	Message   string `json:"message"`
	Attribute string `json:"attribute"`
}

type Resource struct {
	Name             string
	AttributesMap    map[string]interface{}
	rawAttributes    *json.RawMessage
	ResourceIndex    *ResourceIndex
	AssociationIndex *AssociationIndex
	resourceErrors   *[]ResourceError
}

type ResourceId struct {
	Id string `json:"id"`
}

func NewResource(name string, attributesInterface interface{}) *Resource {
	var attributes map[string]interface{}
	var rawAttributes *json.RawMessage

	if assertedAttributes, ok := attributesInterface.(map[string]interface{}); ok {
		attributes = assertedAttributes
	}

	if assertedAttributes, ok := attributesInterface.(*json.RawMessage); ok {
		rawAttributes = assertedAttributes
		json.Unmarshal(*assertedAttributes, &attributes)
	}

	return &Resource{
		Name:          name,
		AttributesMap: attributes,
		rawAttributes: rawAttributes,
	}
}

func (resource *Resource) AddAssociationIndices(resourceIndex *ResourceIndex, associationIndex *AssociationIndex, resourceErrors *[]ResourceError) {
	resource.ResourceIndex = resourceIndex
	resource.AssociationIndex = associationIndex
	resource.resourceErrors = resourceErrors
}

func (resource *Resource) UUID() (string, error) {
	if uuid, ok := resource.AttributesMap["uuid"].(string); ok {
		return uuid, nil
	} else {
		return "", NewUnimatrixError("Unable to retrieve UUID")
	}
}

func (resource *Resource) RawAttributes() (*json.RawMessage, error) {
	if resource.rawAttributes != nil {
		return resource.rawAttributes, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve raw attributes")
	}
}

func (resource *Resource) Attributes() (map[string]interface{}, error) {
	if resource.AttributesMap != nil {
		return resource.AttributesMap, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attributes")
	}
}

func (resource *Resource) AttributeAsString(name string) (string, error) {
	if attribute, ok := resource.AttributesMap[name].(string); ok {
		return attribute, nil
	} else {
		return "", NewUnimatrixError("Unable to retrieve attribute as string")
	}
}

func (resource *Resource) AttributeAsArray(name string) ([]string, error) {
	if attribute, ok := resource.AttributesMap[name].([]string); ok {
		return attribute, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attribute as array")
	}
}

func (resource *Resource) AttributeAsMap(name string) (map[string]interface{}, error) {
	if attribute, ok := resource.AttributesMap[name].(map[string]interface{}); ok {
		for key, value := range attribute {
			if attribute[key], ok = value.(string); !ok {
				return nil, NewUnimatrixError("Unable to retrieve attribute as map with string values")
			}
		}
		return attribute, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attribute as map")
	}
}

func (resource *Resource) SetAttribute(name string, value interface{}) {
	resource.AttributesMap[name] = value
}

func (resource *Resource) Errors() ([]ResourceError, error) {
	if resource.resourceErrors == nil || resource.AssociationIndex == nil {
		return nil, NewUnimatrixError("Unable to retrieve errors")
	}

	var associationIndex = *resource.AssociationIndex
	var resourceErrors = *resource.resourceErrors
	var errors []ResourceError

	if len(associationIndex[resource.Name][resource.AttributesMap["id"].(string)]["errors"]) > 0 {
		for _, error := range resourceErrors {
			for _, errorId := range associationIndex[resource.Name][resource.AttributesMap["id"].(string)]["errors"] {
				if error.Id == errorId {
					errors = append(errors, error)
					break
				}
			}
		}
	}

	return errors, nil
}

func (resource *Resource) Associations() (ResourceAssociations, error) {
	if resource.ResourceIndex == nil || resource.AssociationIndex == nil {
		return nil, NewUnimatrixError("Unable to retrieve associations")
	}

	var resourceIndex = *resource.ResourceIndex
	var associationIndex = *resource.AssociationIndex
	var associations = make(ResourceAssociations)
	var associationsById = associationIndex[resource.Name][resource.AttributesMap["id"].(string)]

	for associationType, ids := range associationsById {
		for _, id := range ids {
			associations[associationType] = append(associations[associationType], resourceIndex[associationType][id])
		}
	}

	return associations, nil
}

func (resource *Resource) Association(name string) ([]Resource, error) {
	if resource.ResourceIndex == nil || resource.AssociationIndex == nil {
		return nil, NewUnimatrixError("Unable to retrieve association")
	}

	var resourceIndex = *resource.ResourceIndex
	var associationIndex = *resource.AssociationIndex
	var association []Resource
	var associationsById = associationIndex[resource.Name][resource.AttributesMap["id"].(string)]

	for _, id := range associationsById[name] {
		association = append(association, resourceIndex[name][id])
	}

	return association, nil
}
