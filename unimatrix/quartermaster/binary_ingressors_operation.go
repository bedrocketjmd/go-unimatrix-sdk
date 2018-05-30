package quartermaster

import "../../unimatrix"

func NewBinaryIngressorsOperation(realm, binaryUUID string) *unimatrix.Operation {
	path := "/realms/" + realm + "/binaries/" + binaryUUID + "/ingressors"
	return unimatrix.NewOperation(path)
}
