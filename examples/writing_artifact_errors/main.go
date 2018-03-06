package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

func main() {
	accessToken := "1234"

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
	_, error := operation.Write("artifacts", []map[string]string{artifact})

	// 403
	fmt.Println(error)

	// get access token
	unimatrix.SetAuthorizationURL("http://us-west-2.keymaker.acceptance.boxxspring.net")
	clientId := os.Getenv("KEYMAKER_CLIENT")
	clientSecret := os.Getenv("KEYMAKER_SECRET")
	accessTokenOperation := unimatrix.NewAccessTokenOperation(clientId, clientSecret)
	tokenResponse, _ := accessTokenOperation.Read()
	accessTokenValid := tokenResponse["access_token"].(string)

	// pass in access token
	operation.AssignParameters(map[string][]string{"access_token": []string{accessTokenValid}})

	artifact["uuid"] = "does-not-exist"

	response, _ := operation.Write("artifacts", []map[string]string{artifact})

	// 200 with attribute errors
	fmt.Println(response[0].GetErrors())
}
