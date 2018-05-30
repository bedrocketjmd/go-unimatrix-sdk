package distributor

import "../../unimatrix"

func NewDistributionsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "distributions")
}
