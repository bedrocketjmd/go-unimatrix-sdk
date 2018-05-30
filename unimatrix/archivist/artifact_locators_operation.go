package archivist

import "../../unimatrix"

func NewArtifactLocatorsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "artifact_locators")
}
