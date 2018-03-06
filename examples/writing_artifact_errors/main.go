package main

import (
	"fmt"

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

	fmt.Println(error)
}
