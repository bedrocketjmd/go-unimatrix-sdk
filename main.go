package main

import (
	"fmt"
	"./unimatrix_sdk"
)

func main() {
	apiUrl := "http://archivist-acceptance-1784742539.us-west-2.elb.amazonaws.com"
	object, err := unimatrix_sdk.Read( apiUrl + "/realms/1e338862026376dd593425404a4f75c0/artifacts" )

	fmt.Println( object.Resources, err )
	fmt.Println( object, err )
}