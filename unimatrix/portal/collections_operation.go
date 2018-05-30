package portal

import "../../unimatrix"

func NewCollectionsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "collections")
}
