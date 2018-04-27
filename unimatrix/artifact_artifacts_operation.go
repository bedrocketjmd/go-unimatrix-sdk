package unimatrix

func NewArtifactArtifactsOperation(realm, artifactUUID string) *Operation {
	path := "/realms/" + realm + "/artifacts/" + artifactUUID + "/artifacts"
	return NewOperation(path)
}
