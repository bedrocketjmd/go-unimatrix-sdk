package unimatrix

func NewDistibutionActivitiesOperation(realm, distributionUUID string) *Operation {
	path := "/realms/" + realm + "/distributions/" + distributionUUID + "/activities"
	return NewOperation(path)
}
