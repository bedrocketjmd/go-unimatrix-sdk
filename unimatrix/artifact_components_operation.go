package unimatrix

func NewArtifactComponentsOperation(realm, artifactUUID string) *Operation {
	path := "/realms/" + realm + "/artifacts/" + artifactUUID + "/components"
	return NewOperation(path)
}
