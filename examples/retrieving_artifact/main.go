package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://archivist-acceptance-1784742539.us-west-2.elb.amazonaws.com")
	operation := unimatrix.NewOperation("/realms/1e338862026376dd593425404a4f75c0/artifacts")
	params := map[string]string{"uuid": "643a9e056300d54eed7a14066513f435"}
	operation.SetParameters(params)
	response, _ := operation.Read()

	for _, resource := range response.Resources {
		fmt.Println(resource)
	}
}
