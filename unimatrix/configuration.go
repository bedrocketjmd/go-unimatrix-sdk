package unimatrix

import (
	"sync"
)

const defaultURL = "http://us-west-2.api.unimatrix.io"
const defaultAuthorizationURL = "http://us-west-2.keymaker.boxxspring.net"

type configuration struct {
	authorizationURL string
	apiURL           string
}

var configurationInstance *configuration
var once sync.Once

func GetConfiguration() *configuration {
	once.Do(func() {
		configurationInstance = &configuration{}
		configurationInstance.apiURL = defaultURL
		configurationInstance.authorizationURL = defaultAuthorizationURL
	})
	return configurationInstance
}

func SetURL(url string) {
	GetConfiguration().apiURL = url
}

func GetURL() string {
	return GetConfiguration().apiURL
}

func SetAuthorizationURL(url string) {
	GetConfiguration().authorizationURL = url
}

func GetAuthorizationURL() string {
	return GetConfiguration().authorizationURL
}
