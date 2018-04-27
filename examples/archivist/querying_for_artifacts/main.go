package main

import (
	"fmt"

	"../../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewRealmOperation(
		"5cbc6bb3db90e2f1236e005f9054776c",
		"artifacts",
	)

	// assign query parameters directly with a map[string][]string
	parameters := map[string][]string{"test": []string{"test_value"}}
	operation.AssignParameters(parameters)
	fmt.Println("OPERATION:")
	fmt.Println(operation)

	// create a new query
	query := unimatrix.NewQuery().
		Where("type_name:eq", "video_artifact").
		WhereArray("name:in", []string{"Final Nick", "test with categoryyy"}).
		WhereArray("relationships.category:eq", []string{"59419fff92047f6e01a98ed35eb21f10"}).
		Where("description:search", "stuff")
	fmt.Println("\nQUERY PARAMETERS:")
	fmt.Println(query.Parameters())

	// assigning query parameters will replace previous query parameters
	operation.AssignParameters(query.Parameters())
	fmt.Println("\nUPDATED OPERATION:")
	fmt.Println(operation)

	// create another query
	additional_query := unimatrix.NewQuery().
		Count(5).
		Offset(3).
		Sort("updated_at", "desc").
		WhereArray("name:in", []string{"Another Field"})
	fmt.Println("\nADDITIONAL QUERY PARAMETERS:")
	fmt.Println(query.Parameters())

	// appending query parameters will add to previous query parameters
	operation.AppendParameters(additional_query.Parameters())
	fmt.Println("\nFINAL OPERATION:")
	fmt.Println(operation)

	response, _ := operation.Read()
	resources, _ := response.Resources()
	fmt.Println("\nRESPONSE:")
	for _, resource := range resources {
		fmt.Println(resource)
	}
}
