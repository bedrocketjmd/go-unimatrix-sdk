package unimatrix

func NewCollectionOperation(realm, collectionID string) *Operation {
	path := "/realms/" + realm + "/collections/" + collectionID
	return NewOperation(path)
}
