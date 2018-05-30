package distributor

import "../../unimatrix"

func NewDistibutionActivitiesOperation(realm, distributionUUID string) *unimatrix.Operation {
	path := "/realms/" + realm + "/distributions/" + distributionUUID + "/activities"
	return unimatrix.NewOperation(path)
}
