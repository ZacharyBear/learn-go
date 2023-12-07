package api

import (
	"learn-wire/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	UserService *service.UserService
}

func (u *UserAPI) Router(router *gin.Engine) {
	g := router.Group("/user")
	g.GET("", GetAll)
}

func GetAll(c *gin.Context) {
	// get from db ...
	user := []map[string]any{
		{
			"id":   1,
			"name": "Zenkie Bear",
		},
	}
	c.JSON(http.StatusOK, user)
}
