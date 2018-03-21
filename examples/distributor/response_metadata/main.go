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

	response, _ := operation.Read()

	fmt.Println("Response Metadata")
	fmt.Println("*****************")
	name, _ := response.Name()
	fmt.Printf("Name: %s\n", name)

	typeName, _ := response.TypeName()
	fmt.Printf("Type Name: %s\n", typeName)

	ids, _ := response.Ids()
	fmt.Printf("Ids: %v\n", ids)

	count, _ := response.Count()
	fmt.Printf("Count: %v\n", count)

	unlimitedCount, _ := response.UnlimitedCount()
	fmt.Printf("Unlimited Count: %v\n", unlimitedCount)

	offset, _ := response.Offset()
	fmt.Printf("Offset: %v\n", offset)
}
