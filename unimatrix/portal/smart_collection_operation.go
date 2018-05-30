package unimatrix

func NewSmartCollectionOperation(realm, smartCollectionID string) *Operation {
	path := "/realms/" + realm + "/smart_collections/" + smartCollectionID
	return NewOperation(path)
}
