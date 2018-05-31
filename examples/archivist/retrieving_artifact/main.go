package main

import (
	"fmt"

	"../../../unimatrix"
	"../../../unimatrix/archivist"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := archivist.NewArtifactsOperation("1e338862026376dd593425404a4f75c0")

	query := unimatrix.NewQuery().Where("uuid", "643a9e056300d54eed7a14066513f435")
	operation.AssignParameters(query.Parameters())
	response, _ := operation.Read()

	resources, _ := response.Resources()

	for _, resource := range resources {
		fmt.Println(resource.Attributes())
		fmt.Println(resource.UUID())
		fmt.Println(resource.AttributeAsString("name"))
		fmt.Println(resource.RawAttributes())
	}
}
