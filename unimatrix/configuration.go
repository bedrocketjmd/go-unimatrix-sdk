package unimatrix

import (
	"sync"
)

type configuration struct {
	authenticationURL string
	apiURL            string
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

func SetAuthenticationURL(url string) *configuration {
	GetConfiguration()
	instance.authenticationURL = url
	return instance
}

func GetAuthenticationURL() string {
	GetConfiguration()
	if instance.authenticationURL != "" {
		return instance.authenticationURL
	} else {
		return "http://us-west-2.keymaker.boxxspring.net"
	}
}
