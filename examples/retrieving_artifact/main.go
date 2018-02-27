package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewOperation("/realms/1e338862026376dd593425404a4f75c0/artifacts")
	query := unimatrix.NewQuery().Where("uuid", "643a9e056300d54eed7a14066513f435")
	operation.AssignParameters(query.Parameters())
	response, _ := operation.Read()

	for _, resource := range response.Resources {
		fmt.Println(resource)
	}
}
