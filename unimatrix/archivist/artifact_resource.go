package archivist

import "../../unimatrix"

type Artifact struct {
	*unimatrix.Resource
}

func (resource *Artifact) Relationships(name string) ([]unimatrix.Resource, error) {
	if resource.ResourceIndex == nil || resource.AssociationIndex == nil {
		return nil, unimatrix.NewUnimatrixError("Unable to retrieve relationships")
	}

	var resourceIndex = *resource.ResourceIndex
	var associationIndex = *resource.AssociationIndex
	var association []unimatrix.Resource
	var associationsById = associationIndex["artifacts"][resource.AttributesMap["id"].(string)]

	for _, id := range associationsById["artifact_relationships"] {
		relationship := resourceIndex["artifact_relationships"][id]
		relationshipName, _ := relationship.AttributeAsString("name")
		if relationshipName == name {
			association = append(association, relationship)
		}
	}
	return association, nil
}

func (resource *Artifact) RelatedArtifacts(name string) ([]unimatrix.Resource, error) {
	if resource.ResourceIndex == nil || resource.AssociationIndex == nil {
		return nil, unimatrix.NewUnimatrixError("Unable to retrieve related artifacts")
	}

	var resourceIndex = *resource.ResourceIndex
	var associationIndex = *resource.AssociationIndex
	var association []unimatrix.Resource
	var associationsById = associationIndex["artifacts"][resource.AttributesMap["id"].(string)]

	for _, id := range associationsById["artifact_relationships"] {
		relationship := resourceIndex["artifact_relationships"][id]
		relationshipName, _ := relationship.AttributeAsString("name")
		if relationshipName == name {
			relationshipId := relationship.AttributesMap["id"].(string)
			relatedId := associationIndex["artifact_relationships"][relationshipId]["artifacts"][0]
			association = append(association, resourceIndex["artifacts"][relatedId])
		}
	}
	return association, nil
}
