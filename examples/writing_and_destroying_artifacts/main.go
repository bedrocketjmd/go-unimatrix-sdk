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
	operation.SetAccessToken(accessToken)

	// create artifacts
	artifact1 := unimatrix.NewResource("artifacts", make(map[string]interface{}), nil, nil)
	artifact1.SetAttribute("type_name", "video_artifact")
	artifact1.SetAttribute("provider", "Boxxspring")
	artifact1.SetAttribute("provider_uid", "go_sdk_test")
	artifact1.SetAttribute("name", "Go SDK Test")

	artifact2 := unimatrix.NewResource("artifacts", make(map[string]interface{}), nil, nil)
	artifact2.SetAttribute("type_name", "video_artifact")
	artifact2.SetAttribute("provider", "Boxxspring")
	artifact2.SetAttribute("provider_uid", "go_sdk_test")
	artifact2.SetAttribute("name", "Go SDK Test 2")

	artifacts := []unimatrix.Resource{*artifact1, *artifact2}

	// write artifacts
	writeResponse, _ := operation.WriteResources("artifacts", artifacts)
	fmt.Println(writeResponse)
	var uuids []string
	for _, artifact := range writeResponse {
		uuid, _ := artifact.GetUUID()
		uuids = append(uuids, uuid)
	}

	// destroy artifact
	destroyResponse, _ := operation.DestroyByUUIDs(uuids)
	fmt.Println(destroyResponse)
}
