package unimatrix

func NewArtifactsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "artifacts")
}
