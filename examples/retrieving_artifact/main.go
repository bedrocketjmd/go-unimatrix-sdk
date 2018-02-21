package main

import (
	"fmt"
	"../../unimatrix"
)

func main() {
	apiUrl := "http://archivist-acceptance-1784742539.us-west-2.elb.amazonaws.com"
	operation := unimatrix.NewOperation(apiUrl + "/realms/1e338862026376dd593425404a4f75c0/artifacts")
	resources, _ := operation.Read()

	for _, resource := range resources {
		fmt.Println(resource)
	}
}
