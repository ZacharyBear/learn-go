package router

import (
	"learn-wire/api"

	"github.com/gin-gonic/gin"
)

type RouterAPI struct {
	UserAPI *api.UserAPI
}

func InitRouter(app *RouterAPI) {
	router := gin.New()

	if app.UserAPI != nil {
		app.UserAPI.Router(router)
	}
	// ... config with router

	router.Run("localhost:8888")
}
