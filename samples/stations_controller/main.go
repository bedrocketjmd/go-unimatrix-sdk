package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid> <station-uuid>
// go run samples/stations_controller/main.go 4327d464be1ef01c81e01dd7c65a2f7e 76f35575c9a239a5e0606d37c0f8dec2

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")

	// Params
	realmUuid := os.Args[1]
	var stationUuid string
	if len(os.Args) > 2 {
		stationUuid = os.Args[2]
	}

	// Query or Read
	operation := unimatrix.NewRealmScopedOperation(realmUuid, "artifacts")
	queryType := "Query"

	query := unimatrix.NewQuery().
		Where("type_name", "station_artifact")

	operation.AssignParameters(query.Parameters())

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
