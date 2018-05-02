package unimatrix

func NewActivitiesOperation(realm string) *Operation {
	return NewRealmOperation(realm, "activities")
}
