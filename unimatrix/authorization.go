package unimatrix

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AccessTokenOperation struct {
	clientId     string
	clientSecret string
}

func NewAccessTokenOperation(clientId string, clientSecret string) *AccessTokenOperation {
	return &AccessTokenOperation{clientId: clientId, clientSecret: clientSecret}
}

func (accessTokenOperation *AccessTokenOperation) AccessToken() (map[string]interface{}, error) {
	client := &http.Client{}

	authorizationURL := GetAuthorizationURL() + "/token"

	req, err := http.NewRequest("POST", authorizationURL, nil)

	if err != nil {
		return nil, err
	}

	query := req.URL.Query()

	query.Add("grant_type", "client_credentials")
	query.Add("client_id", accessTokenOperation.clientId)
	query.Add("client_secret", accessTokenOperation.clientSecret)

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
