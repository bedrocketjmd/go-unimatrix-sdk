package main

import (
	"fmt"

	"../../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewRealmScopedOperation(
		"1e338862026376dd593425404a4f75c0",
		"artifacts",
	)

	query := unimatrix.NewQuery().Where("uuid", "643a9e056300d54eed7a14066513f435")
	operation.AssignParameters(query.Parameters())
	resources, _ := operation.Read()

	for _, resource := range resources {
		fmt.Println(resource.GetAttributes())
		fmt.Println(resource.GetUUID())
		fmt.Println(resource.GetAttributeAsString("name"))
		fmt.Println(resource.GetRawAttributes())
	}
}
