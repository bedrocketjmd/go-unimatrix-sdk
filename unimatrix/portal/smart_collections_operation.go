package portal

import "../../unimatrix"

func NewSmartCollectionsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "smart_collections")
}
