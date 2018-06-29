package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjyi/apiserver/pkgs/errno"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{code, message, data})
}
