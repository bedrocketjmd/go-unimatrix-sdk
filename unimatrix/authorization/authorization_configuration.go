package unimatrix

import (
	"sync"
)

const defaultAuthorizationURL = "http://us-west-2.keymaker.boxxspring.net"

type authorizationConfiguration struct {
	authorizationURL string
}

var authorizationConfigurationInstance *authorizationConfiguration
var authorizationConfigurationSetup sync.Once

func GetAuthorizationConfiguration() *authorizationConfiguration {
	authorizationConfigurationSetup.Do(func() {
		authorizationConfigurationInstance = &authorizationConfiguration{}
		authorizationConfigurationInstance.authorizationURL = defaultAuthorizationURL
	})
	return authorizationConfigurationInstance
}

func SetAuthorizationURL(url string) {
	GetAuthorizationConfiguration().authorizationURL = url
}

func GetAuthorizationURL() string {
	return GetAuthorizationConfiguration().authorizationURL
}
