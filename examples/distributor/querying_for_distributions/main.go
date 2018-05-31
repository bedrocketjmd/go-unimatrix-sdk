package main

import (
	"fmt"
	"os"

	"../../../unimatrix"
	"../../../unimatrix/distributor"
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
	operation := distributor.NewDistributionsOperation("1e338862026376dd593425404a4f75c0")
	operation.SetAccessToken(accessToken)

	// query := unimatrix.NewQuery()

	response, _ := operation.Read()
	resources, _ := response.Resources()
	for _, resource := range resources {
		fmt.Println(resource)
	}
}
