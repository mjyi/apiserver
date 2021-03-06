package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/mjyi/apiserver/handler"
	"github.com/mjyi/apiserver/model"
	"github.com/mjyi/apiserver/pkgs/errno"
)

// Create :create a new account.
func Create(c *gin.Context) {

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		log.Fatal("db error.", err)
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{r.Username}
	handler.SendResponse(c, nil, rsp)
}
