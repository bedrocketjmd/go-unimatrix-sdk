package unimatrix

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func Request(url string, method string, parameters map[string][]string) (*Parser, error) {
	client := &http.Client{}

	req, error := http.NewRequest(method, url, nil)

	if error != nil {
		return nil, error
	}

	req.URL.RawQuery = RawParameters(parameters)

	resp, error := client.Do(req)

	if error != nil {
		return nil, error
	}

	bodyText, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}

	parser := NewParser(bodyText)

	return parser, nil
}

func RawParameters(parameters map[string][]string) string {
	var rawParameters url.Values
	rawParameters = parameters
	return rawParameters.Encode()
}
