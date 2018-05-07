package unimatrix

func NewSmartCollectionsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "smart_collections")
}
