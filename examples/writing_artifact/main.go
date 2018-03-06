package main

import (
	"fmt"
	"os"

	"../../unimatrix"
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

	// pass in access token
	operation.AssignParameters(map[string][]string{"access_token": []string{accessToken}})

	// create artifact
	artifact := map[string]string{}
	artifact["type_name"] = "video_artifact"
	artifact["provider"] = "Boxxspring"
	artifact["provider_uid"] = "go_sdk_test"
	artifact["name"] = "Go SDK Test"

	// write artifact
	response, _ := operation.Write("artifacts", []map[string]string{artifact})
	fmt.Println(response)
}
