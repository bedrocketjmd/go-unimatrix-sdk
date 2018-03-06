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

func (resource *Resource) GetUUID() string {
	return resource.attributes["uuid"].(string)
}

func (resource *Resource) GetAttributes() ResourceAttributes {
	return resource.attributes
}

func (resource *Resource) GetAttributeAsString(name string) string {
	return resource.attributes[name].(string)
}

func (resource *Resource) GetAttributeAsArray(name string) []string {
	return resource.attributes[name].([]string)
}

func (resource *Resource) GetAttributeAsMap(name string) map[string]string {
	return resource.attributes[name].(map[string]string)
}

func (resource *Resource) SetAttribute(name string, value interface{}) {
	resource.attributes[name] = value
}

func (resource *Resource) GetErrors() []ResourceError {
	return resource.errors
}

func (resource *Resource) GetAssociations() ResourceAssociations {
	var associations ResourceAssociations

	for associationType, ids := range resource.associationsById {
		for _, id := range ids {
			associations[associationType] = append(associations[associationType], resourceIndex[resource.name][id])
		}
	}

	return associations
}

func (resource *Resource) GetAssociation(name string) []Resource {
	var association []Resource

	for _, id := range resource.associationsById[name] {
		association = append(association, resourceIndex[name][id])
	}

	return association
}
