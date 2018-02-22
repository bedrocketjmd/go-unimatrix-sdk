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
var once sync.Once

func GetConfiguration() *configuration {
	once.Do(func() {
		configurationInstance = &configuration{}
		configurationInstance.URL = defaultURL
		configurationInstance.authorizationURL = defaultAuthorizationURL
	})
	return configurationInstance
}

func SetURL(url string) {
	GetConfiguration().URL = url
}

func GetURL() string {
	return GetConfiguration().URL
}

func SetAuthorizationURL(url string) {
	GetConfiguration().authorizationURL = url
}

func GetAuthorizationURL() string {
	return GetConfiguration().authorizationURL
}
