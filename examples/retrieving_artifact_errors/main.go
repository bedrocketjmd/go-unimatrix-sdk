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

	query := unimatrix.NewQuery().
		Where("uuid", "does-not-exist")

	operation.AssignParameters(query.Parameters())
	_, error := operation.Read()

	fmt.Println(error)
}
