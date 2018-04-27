package unimatrix

func NewArtifactsArtifactsOperation(realm, artifactUUID string) *Operation {
	path := "/realms/" + realm + "/artifacts/" + artifactUUID + "/artifacts"
	return NewOperation(path)
}
