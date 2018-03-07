package unimatrix

type ResourceAttributes map[string]interface{}

type ResourceAssociations map[string][]Resource

type ResourceError struct {
	Id        string `json:"id"`
	TypeName  string `json:"type_name"`
	Message   string `json:"message"`
	Attribute string `json:"attribute"`
}

type Resource struct {
	name             string
	attributes       ResourceAttributes
	associationsById map[string][]string
	errors           []ResourceError
}

var associations = ResourceAssociations{}

func NewResource(name string, resourceAttributes ResourceAttributes, associationsById map[string][]string, errors []ResourceError) *Resource {
	return &Resource{
		name:             name,
		attributes:       resourceAttributes,
		associationsById: associationsById,
		errors:           errors,
	}
}

func (resource *Resource) GetUUID() (string, error) {
	if uuid, ok := resource.attributes["uuid"].(string); ok {
		return uuid, nil
	} else {
		return "", NewUnimatrixError("Unable to retrieve UUID")
	}
}

func (resource *Resource) GetAttributes() (ResourceAttributes, error) {
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
	if resource.errors != nil {
		return resource.errors, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve errors")
	}
}

func (resource *Resource) GetAssociations() (ResourceAssociations, error) {
	var associations ResourceAssociations

	if resource.associationsById != nil {
		for associationType, ids := range resource.associationsById {
			for _, id := range ids {
				associations[associationType] = append(associations[associationType], resourceIndex[resource.name][id])
			}
		}

		return associations, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve associations")
	}
}

func (resource *Resource) GetAssociation(name string) ([]Resource, error) {
	var association []Resource

	if resource.associationsById[name] != nil {
		for _, id := range resource.associationsById[name] {
			association = append(association, resourceIndex[name][id])
		}

		return association, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve association")
	}
}
