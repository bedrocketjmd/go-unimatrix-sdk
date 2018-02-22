package main

import (
	"fmt"
	"os"

	"../../unimatrix"
)

func main() {
	unimatrix.SetAuthenticationURL("http://us-west-2.keymaker.acceptance.boxxspring.net")

	clientId := os.Getenv("KEYMAKER_CLIENT")
	clientSecret := os.Getenv("KEYMAKER_SECRET")
	clientOperation := unimatrix.NewClientOperation(clientId, clientSecret)

	tokenResponse, _ := clientOperation.AccessToken()

	fmt.Println(tokenResponse)
}
