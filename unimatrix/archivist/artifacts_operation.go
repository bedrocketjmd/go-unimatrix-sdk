package archivist

import "../../unimatrix"

func NewArtifactsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "artifacts")
}
