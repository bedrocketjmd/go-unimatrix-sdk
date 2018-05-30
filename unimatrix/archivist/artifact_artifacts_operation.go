package archivist

import "../../unimatrix"

func NewArtifactArtifactsOperation(realm, artifactUUID string) *unimatrix.Operation {
	path := "/realms/" + realm + "/artifacts/" + artifactUUID + "/artifacts"
	return unimatrix.NewOperation(path)
}
