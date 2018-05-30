package historian

import "../../unimatrix"

func NewHistoriesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "histories")
}
