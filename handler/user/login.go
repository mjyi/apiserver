package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mjyi/apiserver/handler"
	"github.com/mjyi/apiserver/model"
	"github.com/mjyi/apiserver/pkgs/auth"
	"github.com/mjyi/apiserver/pkgs/errno"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	//
	//// Sign the json web token.
	//t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	//if err != nil {
	//	handler.SendResponse(c, errno.ErrToken, nil)
	//	return
	//}
	//
	//handler.SendResponse(c, nil, model.Token{Token: t})

}
