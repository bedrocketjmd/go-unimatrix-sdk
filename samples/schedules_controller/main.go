package main

import (
	"fmt"
	"os"
	"strings"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid> <station-uuids> <start-time-gte>
// go run samples/schedules_controller/main.go 4327d464be1ef01c81e01dd7c65a2f7e c14569f66b4544098cc435bedfa44179,76f35575c9a239a5e0606d37c0f8dec2 2018-02-25T03:00:00Z

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")

	// Params
	realmUuid := os.Args[1]
	station_uuids := strings.Split(os.Args[2], ",")
	var startTimeGte string
	if len(os.Args) > 3 {
		startTimeGte = os.Args[3]
	}

	typeNames := []string{
		"schedule_movie_artifact",
		"schedule_episode_artifact",
		"schedule_competition_artifact",
	}
	// timeParameters := []string{
	// 	"startTime.gt",
	// 	"startTime.gte",
	// 	"startTime.lt",
	// 	"startTime.lte",
	// }

	// Query By Station
	operation := unimatrix.NewRealmScopedOperation(realmUuid, "artifacts")

	query := unimatrix.NewQuery().
		WhereArray("type_name:in", typeNames).
		WhereArray("relationships.category:in", station_uuids).
		Include("relationships.category", "artifacts")

	operation.AssignParameters(query.Parameters())

	if startTimeGte != "" {
		timeQuery := unimatrix.NewQuery().Where("originated_at:gte", startTimeGte)
		operation.AppendParameters(timeQuery.Parameters())
	}

	response, _ := operation.Read()
	resources, _ := response.Resources()

	var stationScheduleEvents = make(map[string][]unimatrix.Resource)

	for _, resource := range resources {
		categories, _ := resource.RelatedArtifacts("category")
		category := categories[0]
		typeName, _ := category.AttributeAsString("type_name")

		if typeName == "station_artifact" {
			uuid, _ := category.UUID()
			var scheduleEvents []unimatrix.Resource

			if _, ok := stationScheduleEvents[uuid]; ok {
				scheduleEvents = stationScheduleEvents[uuid]
			}

			scheduleEvents = append(scheduleEvents, resource)

			stationScheduleEvents[uuid] = scheduleEvents
		}
	}

	fmt.Println("Query Schedules By Stations")
	fmt.Println("*****************")
	fmt.Println(stationScheduleEvents)
}
