package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewOperation("/realms/5cbc6bb3db90e2f1236e005f9054776c/artifacts")
	query := unimatrix.NewQuery().
		Where("type_name:eq", "video_artifact").
		WhereArray("name:in", []string{"Final Nick", "test with categoryyy"}).
		WhereArray("relationships.category:eq", []string{"59419fff92047f6e01a98ed35eb21f10"}).
		Where("description:search", "stuff")
	fmt.Println(query.Parameters())

	operation.SetParameters(query.Parameters())
	response, _ := operation.Read()

	for _, resource := range response.Resources {
		fmt.Println(resource)
	}
}
