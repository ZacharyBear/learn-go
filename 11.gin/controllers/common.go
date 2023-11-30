package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int   `json:"code"`
	Msg   any   `json:"msg"`
	Data  any   `json:"data,omitempty"`
	Count int64 `json:"count,omitempty"`
}

func Success(c *gin.Context, code int, msg any, data any, count int64) {
	json := &Response{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(http.StatusOK, json)
}

func Error(c *gin.Context, code int, msg any) {
	json := &Response{
		Code: code,
		Msg:  msg,
	}
	c.JSON(http.StatusInternalServerError, json)
}
