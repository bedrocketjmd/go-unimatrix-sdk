package unimatrix

func NewBinaryIngressorsOperation(realm, binaryUUID string) *Operation {
	path := "/realms/" + realm + "/binaries/" + binaryUUID + "/ingressors"
	return NewOperation(path)
}
