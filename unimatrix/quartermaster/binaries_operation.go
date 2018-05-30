package quartermaster

import "../../unimatrix"

func NewBinariesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "binaries")
}
