package distributor

import "../../unimatrix"

func NewDestinationActivitiesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "destination_activities")
}
