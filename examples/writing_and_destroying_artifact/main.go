package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

func main() {
	// get access token
	unimatrix.SetAuthorizationURL("http://us-west-2.keymaker.boxxspring.net")
	clientId := os.Getenv("KEYMAKER_CLIENT")
	clientSecret := os.Getenv("KEYMAKER_SECRET")
	accessTokenOperation := unimatrix.NewAccessTokenOperation(clientId, clientSecret)
	tokenResponse, _ := accessTokenOperation.Read()
	accessToken := tokenResponse["access_token"].(string)

	// new operation
	unimatrix.SetURL("http://us-west-2.api.unimatrix.io")
	operation := unimatrix.NewRealmScopedOperation(
		"a5784c49027366bd728b3c24e6bf55c3",
		"artifacts",
	)

	// pass in access token
	operation.AssignParameters(map[string][]string{"access_token": []string{accessToken}})

	// create artifact
	artifact := map[string]string{}
	artifact["type_name"] = "isp_video_artifact"
	artifact["provider"] = "iStreamPlanet"
	artifact["provider_uid"] = "go_sdk_test"
	artifact["name"] = "Go SDK Test"

	// write artifact
	writeResponse, _ := operation.Write("artifacts", []map[string]string{artifact})
	fmt.Println(writeResponse)
	uuid := writeResponse.Resources[0]["uuid"].(string)

	// destroy artifact
	destroyResponse, _ := operation.DestroyByUUID(uuid)
	fmt.Println(destroyResponse)
}
