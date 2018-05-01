package unimatrix

func NewCollectionsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "collections")
}
