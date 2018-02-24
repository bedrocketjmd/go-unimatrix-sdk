package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

func main() {
	unimatrix.SetAuthorizationURL("http://us-west-2.keymaker.acceptance.boxxspring.net")

	clientId := os.Getenv("KEYMAKER_CLIENT")
	clientSecret := os.Getenv("KEYMAKER_SECRET")
	accessTokenOperation := unimatrix.NewAccessTokenOperation(clientId, clientSecret)

	tokenResponse, _ := accessTokenOperation.AccessToken()

	fmt.Println(tokenResponse)
}
