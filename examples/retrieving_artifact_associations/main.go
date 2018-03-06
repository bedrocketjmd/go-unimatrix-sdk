package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewRealmScopedOperation(
		"1e338862026376dd593425404a4f75c0",
		"artifacts",
	)

	query := unimatrix.NewQuery().
		Where("uuid", "4053b259918f6a81b0cac7f2a0e78dcc").
		Include("relationships.category", "artifacts")

	operation.AssignParameters(query.Parameters())
	resources, _ := operation.Read()

	for _, resource := range resources {
		fmt.Println(resource.GetAssociations())
		fmt.Println(resource.GetAssociation("artifact_relationships")[0].GetAssociation("artifacts"))
	}
}
