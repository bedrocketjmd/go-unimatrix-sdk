package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")
	operation := unimatrix.NewRealmScopedOperation(
		"a5784c49027366bd728b3c24e6bf55c3",
		"artifacts",
	)

	// assign query parameters directly with a map[string][]string
	parameters := map[string][]string{"test": []string{"test_value"}}
	operation.AssignParameters(parameters)
	fmt.Println("OPERATION:")
	fmt.Println(operation)

	// create a new query
	query := unimatrix.NewQuery().
		Where("type_name:eq", "isp_video_artifact").
		WhereArray("name:in", []string{"SDK Test isp video artifact", "SDK Test isp video artifact 2"}).
		WhereArray("relationships.category:eq", []string{"072eb7ff34010da7034a562e15a4fea3"}).
		Where("description:search", "Description field test text")
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
	fmt.Println("\nRESPONSE:")
	for _, resource := range response.Resources {
		fmt.Println(resource)
	}
}
