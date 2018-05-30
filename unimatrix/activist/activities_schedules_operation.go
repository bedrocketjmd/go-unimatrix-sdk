package activist

import "../../unimatrix"

func NewActivitiesSchedulesOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "activities_schedules")
}
