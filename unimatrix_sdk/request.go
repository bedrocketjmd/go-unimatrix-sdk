package unimatrix_sdk

import (
	"io/ioutil"
	"net/http"
)

func Request(url string, method string) (*Parser, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	parser := NewParser(bodyText)

	return parser, nil
}
