package fabricator

import "../../unimatrix"

func NewFabricationsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "fabrications")
}
