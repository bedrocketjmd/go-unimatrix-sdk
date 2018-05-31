package main

import (
	"fmt"

	"../../../unimatrix"
	"../../../unimatrix/archivist"
)

func main() {
	operation := archivist.NewArtifactsOperation("1e338862026376dd593425404a4f75c0")

	query := unimatrix.NewQuery().
		Where("uuid", "does-not-exist")

	operation.AssignParameters(query.Parameters())
	_, error := operation.Read()

	fmt.Println(error)
}
