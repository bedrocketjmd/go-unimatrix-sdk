package archivist

import "../../unimatrix"

func NewArtifactComponentsOperation(realm, artifactUUID string) *unimatrix.Operation {
	path := "/realms/" + realm + "/artifacts/" + artifactUUID + "/components"
	return unimatrix.NewOperation(path)
}
