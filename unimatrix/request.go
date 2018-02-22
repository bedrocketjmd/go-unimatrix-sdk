package unimatrix

import (
	"io/ioutil"
	"net/http"
)

func Request(url string, method string) (*Parser, error) {
	client := &http.Client{}

	req, error := http.NewRequest(method, url, nil)

	if error != nil {
		return nil, error
	}

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
