package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid>
//
//
// Query:
// go run samples/genres_controller/main.go 66f3d77a8d522efab771baf740384037

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")

	// Params
	realmUuid := os.Args[1]

	// Query
	operation := unimatrix.NewRealmOperation(realmUuid, "artifacts")

	query := unimatrix.NewQuery().
		Where("type_name", "genre_artifact")

	operation.AssignParameters(query.Parameters())

	response, _ := operation.Read()
	resources, _ := response.Resources()

	fmt.Println("Query Genres")
	fmt.Println("*****************")
	fmt.Println(resources)
}
