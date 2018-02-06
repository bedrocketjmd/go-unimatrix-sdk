package unimatrix_sdk

import (
	"net/http"
	"io/ioutil"
)

func Request( url string, method string ) ( UnimatrixObject, error ) {
	client := &http.Client{}

	req, err := http.NewRequest( method, url, nil )

	if err != nil {
		return UnimatrixObject{}, err
	}

	resp, err := client.Do( req )

	if err != nil {
		return UnimatrixObject{}, err
	}

	bodyText, err := ioutil.ReadAll( resp.Body )

	if err != nil {
		return UnimatrixObject{}, err
	}

	parsedResponse := Parse( bodyText )

	return parsedResponse, nil
}