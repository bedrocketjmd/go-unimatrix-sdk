package main

import (
	"fmt"

	"../../../unimatrix"
	"../../../unimatrix/archivist"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := archivist.NewArtifactsOperation("5cbc6bb3db90e2f1236e005f9054776c")

	// create a new query
	query := unimatrix.NewQuery().
		Where("type_name:eq", "video_artifact").
		Count(5).
		Offset(3)

	// assigning query parameters will replace previous query parameters
	operation.AssignParameters(query.Parameters())

	response, _ := operation.Read()

	fmt.Println("Response Metadata")
	fmt.Println("*****************")
	name, _ := response.Name()
	fmt.Printf("Name: %s\n", name)

	typeName, _ := response.TypeName()
	fmt.Printf("Type Name: %s\n", typeName)

	ids, _ := response.Ids()
	fmt.Printf("Ids: %v\n", ids)

	count, _ := response.Count()
	fmt.Printf("Count: %v\n", count)

	unlimitedCount, _ := response.UnlimitedCount()
	fmt.Printf("Unlimited Count: %v\n", unlimitedCount)

	offset, _ := response.Offset()
	fmt.Printf("Offset: %v\n", offset)
}
