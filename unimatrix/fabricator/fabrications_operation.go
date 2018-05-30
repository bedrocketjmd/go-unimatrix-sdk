package unimatrix

func NewFabricationsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "fabrications")
}
