package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	Success(c, 0, "Success", "{\"name\":\"Zenkie Bear\"}", 1)
}

func GetList(c *gin.Context) {
	Error(c, 4004, "This user is not exists.")
}
