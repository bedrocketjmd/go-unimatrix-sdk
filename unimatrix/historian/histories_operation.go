package unimatrix

func NewHistoriesOperation(realm string) *Operation {
	return NewRealmOperation(realm, "histories")
}
