package main

import (
	"fmt"

	"../../../unimatrix"
)

func main() {
	unimatrix.SetAuthorizationURL("http://us-west-2.keymaker.acceptance.boxxspring.net")

	clientId := "1234"
	clientSecret := "5678"
	accessTokenOperation := unimatrix.NewAccessTokenOperation(clientId, clientSecret)

	_, error := accessTokenOperation.Read()

	fmt.Println(error)
}
