package main

import (
	"fmt"
	"os"
	"time"

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

	// new artifact operation
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	artifactOperation := unimatrix.NewRealmScopedOperation(
		"1e338862026376dd593425404a4f75c0",
		"artifacts",
	)
	artifactOperation.SetAccessToken(accessToken)

	// create artifact for distribution
	artifact := unimatrix.NewResource("artifacts", make(map[string]interface{}))
	artifact.SetAttribute("type_name", "video_artifact")
	artifact.SetAttribute("provider", "Boxxspring")
	artifact.SetAttribute("provider_uid", "go_sdk_test")
	artifact.SetAttribute("name", "Go SDK Test")

	// write artifact
	fmt.Println("Writing artifact")
	artifactWriteResponse, _ := artifactOperation.WriteResource("artifacts", *artifact)
	artifacts, _ := artifactWriteResponse.Resources()
	fmt.Println(artifacts)
	fmt.Println("*****************************************")
	artifact_uuid, _ := artifacts[0].UUID()

	// new distribution operation
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	distributionOperation := unimatrix.NewRealmScopedOperation(
		"1e338862026376dd593425404a4f75c0",
		"distributions",
	)
	distributionOperation.SetAccessToken(accessToken)

	// create distribution
	distribution := unimatrix.NewResource("distributions", make(map[string]interface{}))
	distribution.SetAttribute("type_name", "boxxspring_distribution")
	distribution.SetAttribute("name", "Go SDK Test Distribution")
	distribution.SetAttribute("destination_uuid", "3f4b59f5ede9d6b86538d52466c1c8e8")
	distribution.SetAttribute("artifact_uuid", artifact_uuid)

	// write distribution
	fmt.Println("Writing distribution")
	distributionWriteResponse, _ := distributionOperation.WriteResource("distributions", *distribution)
	distributions, _ := distributionWriteResponse.Resources()
	fmt.Println(distributions)
	fmt.Println("*****************************************")
	distributionUuid, _ := distributions[0].UUID()

	query := unimatrix.NewQuery().
		Where("uuid", distributionUuid)

	distributionOperation.AppendParameters(query.Parameters())

	fmt.Print("Waiting for distribution to complete")

	for i := 0; i < 15; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")

		distributionWriteResponse, _ = distributionOperation.Read()
		distributions, _ = distributionWriteResponse.Resources()
		state, _ := distributions[0].AttributeAsString("state")

		if state == "added" {
			fmt.Println("\nDistribution Completed")
			break
		}
	}

	fmt.Println(distributionWriteResponse)
	fmt.Println("*****************************************")

	// removing distribution
	fmt.Println("Removing distributed artifact")
	destroyResponse, _ := distributionOperation.DestroyByUUID(distributionUuid)
	distributions, _ = destroyResponse.Resources()
	fmt.Println(distributions)
	fmt.Println("*****************************************")

	fmt.Print("Waiting for distributed artifact to be removed")

	for i := 0; i < 15; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")

		destroyResponse, _ = distributionOperation.Read()
		distributions, _ = destroyResponse.Resources()
		state, _ := distributions[0].AttributeAsString("state")

		if state == "removed" {
			fmt.Println("\nDistributed artifact removed")
			break
		}
	}

	fmt.Println(destroyResponse)
}
