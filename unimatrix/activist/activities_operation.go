package activist

import "../../unimatrix"

func NewActivitiesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "activities")
}
