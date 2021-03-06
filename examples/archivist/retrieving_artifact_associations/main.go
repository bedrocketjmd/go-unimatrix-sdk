package main

import (
	"fmt"

	"../../../unimatrix"
)

func main() {
	unimatrix.SetURL("http://us-west-2.api.acceptance.unimatrix.io")
	operation := unimatrix.NewRealmOperation(
		"5cbc6bb3db90e2f1236e005f9054776c",
		"artifacts",
	)

	query := unimatrix.NewQuery().
		Where("uuid", "0b801213aed8e27a14df83b8dc02e16e").
		Include("relationships.athlete", "artifacts").
		Include("relationships.category", "artifacts")

	operation.AssignParameters(query.Parameters())
	response, _ := operation.Read()
	resources, _ := response.Resources()
	artifact := resources[0]

	// get all associations as ResourceAssociations type - map[string][]Resource
	associations, _ := artifact.Associations()
	fmt.Println("ALL ASSOCIATIONS:")
	fmt.Println(associations)

	// get single association
	relationships, _ := artifact.Association("artifact_relationships")
	fmt.Println("\nSINGLE ASSOCIATION:")
	fmt.Println(relationships)

	// get artifact relationships of particular name
	athletes, _ := artifact.Relationships("athlete")
	fmt.Println("\nATHLETE RELATIONSHIPS:")
	fmt.Println(athletes)
	categories, _ := artifact.Relationships("category")
	fmt.Println("\nCATEGORY RELATIONSHIPS:")
	fmt.Println(categories)

	// get related artifact of an artifact relationship
	relationship := relationships[0]
	relatedArtifact, _ := relationship.Association("artifacts")
	fmt.Println("\nRELATED ARTIFACT:")
	fmt.Println(relatedArtifact)

	// get related artifacts of an artifact of particular relationship
	athleteArtifacts, _ := artifact.RelatedArtifacts("athlete")
	categoryArtifacts, _ := artifact.RelatedArtifacts("category")
	fmt.Println("\nRELATED ATHLETE ARTIFACTS:")
	fmt.Println(athleteArtifacts)
	fmt.Println("\nRELATED CATEGORY ARTIFACTS:")
	fmt.Println(categoryArtifacts)
}
