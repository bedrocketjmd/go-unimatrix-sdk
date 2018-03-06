package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")
	operation := unimatrix.NewRealmScopedOperation(
		"a5784c49027366bd728b3c24e6bf55c3",
		"artifacts",
	)

	query := unimatrix.NewQuery().Where("uuid", "040d86e5200117d4c48f9171fe6ede45")
	operation.AssignParameters(query.Parameters())
	response, _ := operation.Read()

	for _, resource := range response.Resources {
		fmt.Println(resource)
	}
}
