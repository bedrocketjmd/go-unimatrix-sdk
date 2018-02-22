package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://archivist-acceptance-1784742539.us-west-2.elb.amazonaws.com")
	operation := unimatrix.NewOperation("/realms/1e338862026376dd593425404a4f75c0/artifacts")

  query := unimatrix.NewQuery().
    Where("uuid", "4053b259918f6a81b0cac7f2a0e78dcc").
    Include("relationships.category", "artifacts")

  operation.SetParameters(query.Parameters())
	response, _ := operation.Read()

	for _, resource := range response.Resources {
		fmt.Println(response.GetAssociations("artifacts", resource["id"]))
	}
}
