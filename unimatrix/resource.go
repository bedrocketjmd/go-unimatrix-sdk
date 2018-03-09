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
	id            string
	name          string
	attributes    map[string]interface{}
	rawAttributes *json.RawMessage
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
		name:          name,
		attributes:    attributes,
		rawAttributes: rawAttributes,
	}
}

func (resource *Resource) GetUUID() (string, error) {
	if uuid, ok := resource.attributes["uuid"].(string); ok {
		return uuid, nil
	} else {
		return "", NewUnimatrixError("Unable to retrieve UUID")
	}
}

func (resource *Resource) GetRawAttributes() (*json.RawMessage, error) {
	if resource.rawAttributes != nil {
		return resource.rawAttributes, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve raw attributes")
	}
}

func (resource *Resource) GetAttributes() (map[string]interface{}, error) {
	if resource.attributes != nil {
		return resource.attributes, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attributes")
	}
}

func (resource *Resource) GetAttributeAsString(name string) (string, error) {
	if attribute, ok := resource.attributes[name].(string); ok {
		return attribute, nil
	} else {
		return "", NewUnimatrixError("Unable to retrieve attribute as string")
	}
}

func (resource *Resource) GetAttributeAsArray(name string) ([]string, error) {
	if attribute, ok := resource.attributes[name].([]string); ok {
		return attribute, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attribute as array")
	}
}

func (resource *Resource) GetAttributeAsMap(name string) (map[string]string, error) {
	if attribute, ok := resource.attributes[name].(map[string]string); ok {
		return attribute, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve attribute as map")
	}
}

func (resource *Resource) SetAttribute(name string, value interface{}) {
	resource.attributes[name] = value
}

func (resource *Resource) GetErrors() ([]ResourceError, error) {
	var errors []ResourceError

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

	if errors != nil {
		return errors, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve errors")
	}
}

func (resource *Resource) GetAssociations() (ResourceAssociations, error) {
	var associations = make(ResourceAssociations)
	var associationsById = associationIndex[resource.name][resource.attributes["id"].(string)]

	if associationsById != nil {
		for associationType, ids := range associationsById {
			for _, id := range ids {
				associations[associationType] = append(associations[associationType], resourceIndex[associationType][id])
			}
		}

		return associations, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve associations")
	}
}

func (resource *Resource) GetAssociation(name string) ([]Resource, error) {
	var association []Resource
	var associationsById = associationIndex[resource.name][resource.attributes["id"].(string)]

	if associationsById[name] != nil {
		for _, id := range associationsById[name] {
			association = append(association, resourceIndex[name][id])
		}

		return association, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve association")
	}
}
