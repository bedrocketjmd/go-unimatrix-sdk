package unimatrix

func NewBlueprintAttributesOperation(realm string) *Operation {
	return NewRealmOperation(realm, "blueprint_attributes")
}
