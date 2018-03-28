package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid>
// go run samples/genres_controller/main.go 4327d464be1ef01c81e01dd7c65a2f7e

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")

	// Params
	realmUuid := os.Args[1]

	// Query
	operation := unimatrix.NewRealmScopedOperation(realmUuid, "artifacts")

	query := unimatrix.NewQuery().
		Where("type_name", "genre_artifact")

	operation.AssignParameters(query.Parameters())

	response, _ := operation.Read()
	resources, _ := response.Resources()

	fmt.Println("Query Genres")
	fmt.Println("*****************")
	fmt.Println(resources)
}
