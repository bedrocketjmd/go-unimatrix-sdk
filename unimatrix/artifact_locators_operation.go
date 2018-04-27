package unimatrix

func NewArtifactLocatorsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "artifact_locators")
}
