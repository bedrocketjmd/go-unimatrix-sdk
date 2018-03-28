package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid> <asset-uuid>
// go run samples/assets_controller/main.go 1e338862026376dd593425404a4f75c0 ff4cd4840eba37f2e2e63fdf7cba2f11

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")

	// Params
	realmUuid := os.Args[1]
	assetUuid := os.Args[2]

	// Read
	operation := unimatrix.NewRealmScopedOperation(realmUuid, "artifacts")

	query := unimatrix.NewQuery().
		Where("type_name", "asset_artifact").
		Where("uuid", assetUuid)

	operation.AssignParameters(query.Parameters())

	response, _ := operation.Read()
	resources, _ := response.Resources()

	fmt.Println("Read Asset UUID:", assetUuid)
	fmt.Println("*****************")
	fmt.Println(resources[0])
}
