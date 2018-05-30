package archivist

import "../../unimatrix"

func NewBlueprintsOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "blueprints")
}
