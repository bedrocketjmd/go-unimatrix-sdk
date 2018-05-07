package unimatrix

func NewActivitiesSchedulesOperation(realm string) *Operation {
	return NewRealmOperation(realm, "activities_schedules")
}
