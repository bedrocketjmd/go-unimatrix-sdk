package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewOperation("/realms/1e338862026376dd593425404a4f75c0/artifacts")

	query := unimatrix.NewQuery().
		Where("uuid", "does-not-exist")

	operation.AssignParameters(query.Parameters())
	_, error := operation.Read()

	fmt.Println(error)
}
