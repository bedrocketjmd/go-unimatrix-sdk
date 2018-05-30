package archivist

import "../../unimatrix"

func NewBlueprintAttributesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "blueprint_attributes")
}
