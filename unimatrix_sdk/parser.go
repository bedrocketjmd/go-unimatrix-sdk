package unimatrix_sdk

import (
	"encoding/json"
)

type Artifact struct {
	uuid string
	realmUuid string
	name string
}

type UnimatrixObject struct {
	Name         string
	TypeName     string
	Keys         []interface{}
	Resources    []Artifact
}

func Parse( rawResponse []byte ) ( UnimatrixObject ) {
	var response map[string]interface{}

	json.Unmarshal([]byte(rawResponse), &response)

	// do something with response

	this := response["$this"].(map[string]interface{})
	name := this["name"].(string)
	artifacts := []Artifact{}

	for _, resource := range response[name].([]interface{}) {
		resource := resource.(map[string]interface{})
		artifact := Artifact{
			uuid: resource["uuid"].(string),
			realmUuid: resource["realm_uuid"].(string),
			name: resource["name"].(string),
		}
		artifacts = append(artifacts, artifact)
	}

	parsedResponse := UnimatrixObject{
		Name: name,
		TypeName: this["type_name"].(string),
		Keys: this["ids"].([]interface{}),
		Resources: artifacts,
	}

	return parsedResponse
}