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
	response, _ := operation.Read()

	for _, resource := range response.Resources {
		fmt.Println(resource)
		fmt.Println(response.GetAssociations("artifacts", resource["id"].(string)))
	}
}
