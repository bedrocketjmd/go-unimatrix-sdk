package unimatrix

func NewBlueprintsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "blueprints")
}
