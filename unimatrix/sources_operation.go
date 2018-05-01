package unimatrix

func NewSourcesOperation(realm string) *Operation {
	return NewRealmOperation(realm, "sources")
}
