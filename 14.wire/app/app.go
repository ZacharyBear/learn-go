package app

import (
	"learn-wire/api"
	"learn-wire/dao/user"
	"learn-wire/router"
	"learn-wire/service"
)

func LoadApp() {
	// create istances

	r := router.RouterAPI{
		UserAPI: &api.UserAPI{
			UserService: &service.UserService{
				UserRepo: &user.UserRepo{},
			},
		},
	}

	router.InitRouter(&r)

}
