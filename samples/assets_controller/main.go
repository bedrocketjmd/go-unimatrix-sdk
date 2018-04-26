package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid> <asset-uuid>
//
//
// Read:
// go run samples/assets_controller/main.go 66f3d77a8d522efab771baf740384037 f7d8bd902ecf36eb04ddd90424d9673d

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")

	// Params
	realmUuid := os.Args[1]
	assetUuid := os.Args[2]

	// Read
	operation := unimatrix.NewRealmOperation(realmUuid, "artifacts")

	query := unimatrix.NewQuery().
		Where("type_name", "asset_artifact").
		Where("uuid", assetUuid)

	operation.AssignParameters(query.Parameters())

	response, _ := operation.Read()
	resources, _ := response.Resources()

	fmt.Println("Read Asset UUID:", assetUuid)
	fmt.Println("*****************")
	fmt.Println(resources)
}
