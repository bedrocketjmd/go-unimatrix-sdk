package unimatrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UnimatrixError struct {
	errorMessage string
	errorStatus  string
	errorCode    int
}

type ErrorResponse struct {
	This struct {
		Name     string  `json:"name"`
		TypeName string  `json:"type_name"`
		Ids      []int64 `json:"ids"`
	} `json:"$this"`
	Errors []struct {
		ID       int64  `json:"id"`
		TypeName string `json:"type_name"`
		Message  string `json:"message"`
	} `json:"errors"`
}

func (unimatrixError *UnimatrixError) Error() string {
	return fmt.Sprintln(unimatrixError.errorMessage)
}

func NewUnimatrixError(err interface{}) error {
	unimatrixError := UnimatrixError{}

	if response, ok := err.(*http.Response); ok {
		unimatrixError.errorStatus = response.Status
		unimatrixError.errorCode = response.StatusCode

		if response.StatusCode == 500 {
			unimatrixError.errorMessage = response.Status + ": An unexpected error occurred."
			return &unimatrixError
		}

		bodyText, error := ioutil.ReadAll(response.Body)

		if error != nil {
			unimatrixError.errorMessage = response.Status + ": An unexpected error occurred."
			return &unimatrixError
		}

		errorResponse := ErrorResponse{}

		error = json.Unmarshal([]byte(bodyText), &errorResponse)

		if error != nil {
			unimatrixError.errorMessage = response.Status + ": An unexpected error occurred."
			return &unimatrixError
		}

		unimatrixError.errorMessage = response.Status + ": " + errorResponse.Errors[0].Message
	} else if e, ok := err.(error); ok {
		unimatrixError.errorMessage = e.Error()
	} else {
		unimatrixError.errorMessage = "An unexpected error occurred."
	}

	return &unimatrixError
}
