package main

import (
	"fmt"
	"os"
	"strings"

	"../../unimatrix"
)

// Test Command
// go run <path-to-controller> <realm-uuid> <station-uuids> <start-time>
//
//
// Query Filtered By Single Station:
// go run samples/schedules_controller/main.go 66f3d77a8d522efab771baf740384037 36addaa58c744a85d8d9fecfa8e17750
//
// Query Filtered By Multiple Stations:
// go run samples/schedules_controller/main.go 66f3d77a8d522efab771baf740384037 36addaa58c744a85d8d9fecfa8e17750,eebde4cc8e8070849bcf172bae5f21f3
//
// Query Filtered By Multiple Stations and Start Time:
// go run samples/schedules_controller/main.go 66f3d77a8d522efab771baf740384037 36addaa58c744a85d8d9fecfa8e17750,eebde4cc8e8070849bcf172bae5f21f3 startTime.gte=2018-02-28T00:23:00Z

func main() {
	// Environment
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")

	// Params
	realmUuid := os.Args[1]
	stationUuids := strings.Split(os.Args[2], ",")
	var startTime string
	if len(os.Args) > 3 {
		startTime = os.Args[3]
	}

	typeNames := []string{
		"schedule_movie_artifact",
		"schedule_episode_artifact",
		"schedule_competition_artifact",
	}

	// Query
	operation := unimatrix.NewRealmScopedOperation(realmUuid, "artifacts")

	// - Filtered by Stations
	query := unimatrix.NewQuery().
		WhereArray("type_name:in", typeNames).
		WhereArray("relationships.category:in", stationUuids).
		Include("relationships.category", "artifacts")

	operation.AssignParameters(query.Parameters())

	// - Filtered by Start Time
	if startTime != "" {
		startTimeSplit := strings.Split(startTime, "=")
		startTimeValue := startTimeSplit[1]
		startTimeOperator := strings.Split(startTimeSplit[0], ".")[1]
		startTimeParam := "originated_at:" + startTimeOperator
		startTimeQuery := unimatrix.NewQuery().
			Where(startTimeParam, startTimeValue)
		operation.AppendParameters(startTimeQuery.Parameters())
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
