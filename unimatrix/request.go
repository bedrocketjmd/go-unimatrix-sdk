package unimatrix

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Request(url string, method string, parameters map[string][]string, body interface{}) (*Response, error) {
	client := &http.Client{}

	requestBody, error := RequestBody(body)

	if error != nil {
		return nil, NewUnimatrixError(error)
	}

	req, error := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if error != nil {
		return nil, NewUnimatrixError(error)
	}

	req.Header.Add("Content-Type", "application/json")
	req.URL.RawQuery = RequestParameters(parameters)

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

	response, error := NewResponse(bodyText)

	if error != nil {
		return nil, error
	}

	return response, nil
}

func RequestParameters(parameters map[string][]string) string {
	var requestParameters url.Values
	requestParameters = parameters
	return requestParameters.Encode()
}

func RequestBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	} else {
		requestBody, error := json.Marshal(body)
		return requestBody, error
	}
}
