package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mjyi/apiserver/handler"
	"github.com/mjyi/apiserver/model"
	"github.com/mjyi/apiserver/pkgs/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
