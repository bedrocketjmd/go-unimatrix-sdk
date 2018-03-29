package main

import (
	"fmt"
	"os"
	"strings"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid> <resource-type> <start-time> <name> <show-name> <call-sign>
// go run samples/search_controller/main.go 4327d464be1ef01c81e01dd7c65a2f7e null startTime.gte=2018-02-25T03:00:00Z testing "Amy is Testing" "Amy is Testing"

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")

	// Constants
	var typeNames = []string{
		"movie_artifact",
		"show_artifact",
		"show_season_artifact",
		"show_episode_artifact",
		"schedule_movie_artifact",
		"schedule_episode_artifact",
		"schedule_competition_artifact",
	}
	var typeNamesMap = make(map[string]struct{})
	for _, typeName := range typeNames {
		typeNamesMap[typeName] = struct{}{}
	}
	var scheduleTypeNames = []string{
		"schedule_movie_artifact",
		"schedule_episode_artifact",
		"schedule_competition_artifact",
	}

	// Params
	realmUuid := os.Args[1]
	var resourceType string
	if len(os.Args) > 2 {
		resourceType = os.Args[2]
	}
	var startTime string
	if len(os.Args) > 3 {
		startTime = os.Args[3]
	}
	var name string
	if len(os.Args) > 4 {
		name = os.Args[4]
	}
	var showName string
	if len(os.Args) > 5 {
		showName = os.Args[5]
	}
	var callSign string
	if len(os.Args) > 6 {
		callSign = os.Args[6]
	}

	// Search Query
	operation := unimatrix.NewRealmScopedOperation(realmUuid, "artifacts")

	resourceTypePresent := false

	// Type Names
	typeNamesToQuery := typeNames

	if resourceType != "" {
		artifactTypeName := resourceType + "_artifact"

		if _, ok := typeNamesMap[artifactTypeName]; ok {
			resourceTypePresent = true
			typeNamesToQuery = []string{artifactTypeName}
		}
	}

	query := unimatrix.NewQuery().
		Include("relationships.category", "artifacts")
	operation.AssignParameters(query.Parameters())

	typeNamesQuery := unimatrix.NewQuery().
		WhereArray("type_name:in", typeNamesToQuery)

	operation.AppendParameters(typeNamesQuery.Parameters())

	// Start Time
	if startTime != "" {
		if !resourceTypePresent {
			typeNamesQuery = unimatrix.NewQuery().
				WhereArray("type_name:in", scheduleTypeNames)
			operation.AppendParameters(typeNamesQuery.Parameters())
		}

		startTimeSplit := strings.Split(startTime, "=")
		startTimeValue := startTimeSplit[1]
		startTimeOperator := strings.Split(startTimeSplit[0], ".")[1]
		startTimeParam := "originated_at:" + startTimeOperator
		startTimeQuery := unimatrix.NewQuery().
			Where(startTimeParam, startTimeValue)
		operation.AppendParameters(startTimeQuery.Parameters())
	}

	// Name
	if name != "" {
		nameQuery := unimatrix.NewQuery().
			Where("name:search", name)
		operation.AppendParameters(nameQuery.Parameters())
	}

	// Show Name
	if showName != "" {
		showNameQuery := unimatrix.NewQuery().
			Where("name", showName).
			Where("relationships.ctegory", "show_artifact")
		operation.AppendParameters(showNameQuery.Parameters())
	}

	// Call Sign
	if callSign != "" {
		callSignQuery := unimatrix.NewQuery().
			Where("short_name", callSign)
		operation.AppendParameters(callSignQuery.Parameters())
	}

	response, _ := operation.Read()
	resources, _ := response.Resources()

	fmt.Println("Search Query")
	fmt.Println("*****************")
	fmt.Println(resources)
}
