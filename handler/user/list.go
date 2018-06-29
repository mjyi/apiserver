package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mjyi/apiserver/handler"
	"github.com/mjyi/apiserver/pkgs/errno"
	"github.com/mjyi/apiserver/service"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{count,infos})

}
