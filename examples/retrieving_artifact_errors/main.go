package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	operation := unimatrix.NewRealmScopedOperation(
		"1e338862026376dd593425404a4f75c0",
		"artifacts",
	)

	query := unimatrix.NewQuery().
		Where("uuid", "does-not-exist")

	operation.AssignParameters(query.Parameters())
	_, error := operation.Read()

	fmt.Println(error)
}
