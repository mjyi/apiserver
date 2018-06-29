package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mjyi/apiserver/handler"
	"github.com/mjyi/apiserver/model"
	"github.com/mjyi/apiserver/pkgs/errno"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
