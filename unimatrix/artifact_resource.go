package unimatrix

func (resource *Resource) GetRelationships(name string) ([]Resource, error) {
	var association []Resource
	var associationsById = associationIndex["artifacts"][resource.attributes["id"].(string)]

	if associationsById["artifact_relationships"] != nil {
		for _, id := range associationsById["artifact_relationships"] {
			relationship := resourceIndex["artifact_relationships"][id]
			relationshipName, _ := relationship.GetAttributeAsString("name")
			if relationshipName == name {
				association = append(association, relationship)
			}
		}
		return association, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve " + name + " relationships")
	}
}

func (resource *Resource) GetRelatedArtifacts(name string) ([]Resource, error) {
	var association []Resource
	var associationsById = associationIndex["artifacts"][resource.attributes["id"].(string)]

	if associationsById["artifact_relationships"] != nil {
		for _, id := range associationsById["artifact_relationships"] {
			relationship := resourceIndex["artifact_relationships"][id]
			relationshipName, _ := relationship.GetAttributeAsString("name")
			if relationshipName == name {
				relationshipId := relationship.attributes["id"].(string)
				relatedId := associationIndex["artifact_relationships"][relationshipId]["artifacts"][0]
				association = append(association, resourceIndex["artifacts"][relatedId])
			}
		}
		return association, nil
	} else {
		return nil, NewUnimatrixError("Unable to retrieve related " + name + " artifacts")
	}
}
