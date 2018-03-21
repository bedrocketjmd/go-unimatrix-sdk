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
		"artifacts",
	)
	operation.SetAccessToken(accessToken)

	// create artifact
	artifact := unimatrix.NewResource("artifacts", make(map[string]interface{}))
	artifact.SetAttribute("type_name", "video_artifact")
	artifact.SetAttribute("provider", "Boxxspring")
	artifact.SetAttribute("provider_uid", "go_sdk_test")
	artifact.SetAttribute("name", "Go SDK Test")

	// write artifact
	writeResponse, _ := operation.WriteResource("artifacts", *artifact)
	resources, _ := writeResponse.Resources()
	fmt.Println(resources)
	uuid, _ := resources[0].UUID()

	// destroy artifact
	destroyResponse, _ := operation.DestroyByUUID(uuid)
	fmt.Println(destroyResponse)
}
