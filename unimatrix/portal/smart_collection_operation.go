package portal

import "../../unimatrix"

func NewSmartCollectionOperation(realm, smartCollectionID string) *unimatrix.Operation {
	path := "/realms/" + realm + "/smart_collections/" + smartCollectionID
	return unimatrix.NewOperation(path)
}
