package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

// Test Commands
// go run <path-to-controller> <realm-uuid> <station-uuid>
//
//
// Query:
// go run samples/stations_controller/main.go 66f3d77a8d522efab771baf740384037
//
// Read:
// go run samples/stations_controller/main.go 66f3d77a8d522efab771baf740384037 36addaa58c744a85d8d9fecfa8e17750

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")

	// Params
	realmUuid := os.Args[1]
	var stationUuid string
	if len(os.Args) > 2 {
		stationUuid = os.Args[2]
	}

	// Query
	operation := unimatrix.NewRealmOperation(realmUuid, "artifacts")
	queryType := "Query"

	query := unimatrix.NewQuery().
		Where("type_name", "station_artifact")

	operation.AssignParameters(query.Parameters())

	// Read
	if stationUuid != "" {
		queryType = "Read"
		readQuery := unimatrix.NewQuery().Where("uuid", stationUuid)
		operation.AppendParameters(readQuery.Parameters())
	}

	response, _ := operation.Read()
	resources, _ := response.Resources()

	fmt.Println(queryType, "Stations")
	fmt.Println("*****************")
	fmt.Println(resources)
}
