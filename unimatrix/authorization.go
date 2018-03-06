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

func (accessTokenOperation *AccessTokenOperation) Read() (map[string]interface{}, error) {
	client := &http.Client{}

	authorizationURL := GetAuthorizationURL() + "/token"

	req, error := http.NewRequest("POST", authorizationURL, nil)

	if error != nil {
		return nil, NewUnimatrixError(error)
	}

	query := req.URL.Query()

	query.Add("grant_type", "client_credentials")
	query.Add("client_id", accessTokenOperation.clientId)
	query.Add("client_secret", accessTokenOperation.clientSecret)

	req.URL.RawQuery = query.Encode()

	resp, error := client.Do(req)

	if error != nil {
		return nil, NewUnimatrixError(error)
	}

	if resp.StatusCode != 200 {
		return nil, NewUnimatrixError(resp)
	}

	bodyText, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		return nil, NewUnimatrixError(error)
	}

	var tokenResponse map[string]interface{}

	error = json.Unmarshal([]byte(bodyText), &tokenResponse)

	if error != nil {
		return nil, NewUnimatrixError(error)
	}

	if tokenResponse["error"] != nil {
		err := UnimatrixError{}

		err.errorMessage = tokenResponse["error"].(string) + ": " + tokenResponse["error_description"].(string)
		err.errorStatus = tokenResponse["error"].(string)

		return nil, &err
	}

	return tokenResponse, nil
}
