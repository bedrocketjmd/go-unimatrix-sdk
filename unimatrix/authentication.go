package unimatrix

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ClientOperation struct {
	clientId     string
	clientSecret string
}

func NewClientOperation(clientId string, clientSecret string) *ClientOperation {
	return &ClientOperation{clientId: clientId, clientSecret: clientSecret}
}

func (clientOperation *ClientOperation) AccessToken() (map[string]interface{}, error) {
	client := &http.Client{}

	authenticationURL := GetAuthenticationURL() + "/token"

	req, err := http.NewRequest("POST", authenticationURL, nil)

	if err != nil {
		return nil, err
	}

	query := req.URL.Query()

	query.Add("grant_type", "client_credentials")
	query.Add("client_id", clientOperation.clientId)
	query.Add("client_secret", clientOperation.clientSecret)

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var tokenResponse map[string]interface{}

	json.Unmarshal([]byte(bodyText), &tokenResponse)

	return tokenResponse, nil
}
