package regent

import "../../unimatrix"

func NewRealmsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "realms")
}
