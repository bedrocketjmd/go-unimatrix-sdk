package curator

import "../../unimatrix"

func NewSourcesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "sources")
}
