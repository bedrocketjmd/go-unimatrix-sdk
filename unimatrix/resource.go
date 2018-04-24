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
	name             string
	attributes       map[string]interface{}
	rawAttributes    *json.RawMessage
	resourceIndex    *ResourceIndex
	associationIndex *AssociationIndex
	resourceErrors   *[]ResourceError
}

type ResourceId struct {
	Id string `json:"id"`
}

func NewResource(name string, attributesInterface interface{}, resourceIndex *ResourceIndex, associationIndex *AssociationIndex, resourceErrors *[]ResourceError) *Resource {
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
		name:             name,
		attributes:       attributes,
		rawAttributes:    rawAttributes,
		resourceIndex:    resourceIndex,
		associationIndex: associationIndex,
		resourceErrors:   resourceErrors,
	}
}

func (resource *Resource) UUID() (string, error) {
	if uuid, ok := resource.attributes["uuid"].(string); ok {
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
	if resource.attributes != nil {
		return resource.attributes, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attributes")
	}
}

func (resource *Resource) AttributeAsString(name string) (string, error) {
	if attribute, ok := resource.attributes[name].(string); ok {
		return attribute, nil
	} else {
		return "", NewUnimatrixError("Unable to retrieve attribute as string")
	}
}

func (resource *Resource) AttributeAsArray(name string) ([]string, error) {
	if attribute, ok := resource.attributes[name].([]string); ok {
		return attribute, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attribute as array")
	}
}

func (resource *Resource) AttributeAsMap(name string) (map[string]string, error) {
	if attribute, ok := resource.attributes[name].(map[string]string); ok {
		return attribute, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attribute as map")
	}
}

func (resource *Resource) SetAttribute(name string, value interface{}) {
	resource.attributes[name] = value
}

func (resource *Resource) Errors() ([]ResourceError, error) {
	var errors []ResourceError
	var associationIndex = *resource.associationIndex
	var resourceErrors = *resource.resourceErrors

	if len(associationIndex[resource.name][resource.attributes["id"].(string)]["errors"]) > 0 {
		for _, error := range resourceErrors {
			for _, errorId := range associationIndex[resource.name][resource.attributes["id"].(string)]["errors"] {
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
	var resourceIndex = *resource.resourceIndex
	var associationIndex = *resource.associationIndex
	var associations = make(ResourceAssociations)
	var associationsById = associationIndex[resource.name][resource.attributes["id"].(string)]

	for associationType, ids := range associationsById {
		for _, id := range ids {
			associations[associationType] = append(associations[associationType], resourceIndex[associationType][id])
		}
	}

	return associations, nil
}

func (resource *Resource) Association(name string) ([]Resource, error) {
	var resourceIndex = *resource.resourceIndex
	var associationIndex = *resource.associationIndex
	var association []Resource
	var associationsById = associationIndex[resource.name][resource.attributes["id"].(string)]

	for _, id := range associationsById[name] {
		association = append(association, resourceIndex[name][id])
	}

	return association, nil
}
