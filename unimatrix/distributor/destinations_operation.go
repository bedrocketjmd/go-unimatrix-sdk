package distributor

import "../../unimatrix"

func NewDestinationsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "destinations")
}
