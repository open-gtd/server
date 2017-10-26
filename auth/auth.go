package auth

import "github.com/open-gtd/server/auth/api"

var apiInitializer = api.Initialize

func Initialize() {
	apiInitializer()
}
