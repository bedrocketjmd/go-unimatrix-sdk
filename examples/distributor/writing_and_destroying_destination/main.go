package main

import (
	"fmt"
	"os"

	"../../../unimatrix"
)

func main() {
	// get access token
	unimatrix.SetAuthorizationURL("http://us-west-2.keymaker.acceptance.boxxspring.net")
	clientId := os.Getenv("KEYMAKER_CLIENT")
	clientSecret := os.Getenv("KEYMAKER_SECRET")
	accessTokenOperation := unimatrix.NewAccessTokenOperation(clientId, clientSecret)
	tokenResponse, _ := accessTokenOperation.Read()
	accessToken := tokenResponse["access_token"].(string)

	// new operation
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewRealmScopedOperation(
		"1e338862026376dd593425404a4f75c0",
		"destinations",
	)
	operation.SetAccessToken(accessToken)

	// create destination
	destination := unimatrix.NewResource("destinations", make(map[string]interface{}))
	destination.SetAttribute("type_name", "boxxspring_destination")
	destination.SetAttribute("name", "Go SDK Test Destination")
	destination.SetAttribute("destination_realm_uuid", "5cbc6bb3db90e2f1236e005f9054776c")

	// write destination
	writeResponse, _ := operation.WriteResource("destinations", *destination)
	fmt.Println(writeResponse)
	uuid, _ := writeResponse[0].UUID()

	// destroy destination
	destroyResponse, _ := operation.DestroyByUUID(uuid)
	fmt.Println(destroyResponse)
}
