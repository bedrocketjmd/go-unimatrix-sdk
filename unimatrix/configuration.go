package unimatrix

import (
	"sync"
)

const defaultURL = "http://us-west-2.api.unimatrix.io"
const defaultAuthorizationURL = "http://us-west-2.keymaker.boxxspring.net"

type configuration struct {
	authorizationURL string
	URL              string
}

var configurationInstance *configuration
var configurationSetup sync.Once

func Configuration() *configuration {
	configurationSetup.Do(func() {
		configurationInstance = &configuration{}
		configurationInstance.URL = defaultURL
		configurationInstance.authorizationURL = defaultAuthorizationURL
	})
	return configurationInstance
}

func SetURL(url string) {
	Configuration().URL = url
}

func URL() string {
	return Configuration().URL
}

func SetAuthorizationURL(url string) {
	Configuration().authorizationURL = url
}

func AuthorizationURL() string {
	return Configuration().authorizationURL
}
