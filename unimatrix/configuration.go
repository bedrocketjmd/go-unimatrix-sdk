package unimatrix

import (
	"sync"
)

type configuration struct {
	authorizationURL string
	apiURL           string
}

var instance *configuration
var once sync.Once

func GetConfiguration() *configuration {
	once.Do(func() {
		instance = &configuration{}
	})
	return instance
}

func SetURL(url string) *configuration {
	GetConfiguration()
	instance.apiURL = url
	return instance
}

func GetURL() string {
	GetConfiguration()
	if instance.apiURL != "" {
		return instance.apiURL
	} else {
		return "http://us-west-2.api.unimatrix.io"
	}
}

func SetAuthorizationURL(url string) *configuration {
	GetConfiguration()
	instance.authorizationURL = url
	return instance
}

func GetAuthorizationURL() string {
	GetConfiguration()
	if instance.authorizationURL != "" {
		return instance.authorizationURL
	} else {
		return "http://us-west-2.keymaker.boxxspring.net"
	}
}
