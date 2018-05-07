package unimatrix

func NewDestinationsOperation(realm string) *Operation {
	return NewRealmOperation(realm, "destinations")
}
