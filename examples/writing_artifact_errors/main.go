package main

import (
	"fmt"

	"../../unimatrix"
)

func main() {
	accessToken := "1234"

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
	artifact["type_name"] = "video_artifact"
	artifact["provider"] = "Boxxspring"
	artifact["provider_uid"] = "go_sdk_test"
	artifact["name"] = "Go SDK Test"

	// write artifact
	_, error := operation.Write("artifacts", []map[string]string{artifact})

	fmt.Println(error)
}
