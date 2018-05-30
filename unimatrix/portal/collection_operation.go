package portal

import "../../unimatrix"

func NewCollectionOperation(realm, collectionID string) *unimatrix.Operation {
	path := "/realms/" + realm + "/collections/" + collectionID
	return unimatrix.NewOperation(path)
}
