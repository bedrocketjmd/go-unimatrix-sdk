package unimatrix

func NewDestinationActivitiesOperation(realm string) *Operation {
	return NewRealmOperation(realm, "destination_activities")
}
