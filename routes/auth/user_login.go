package auth

import (
	"github.com/auula/gws"
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

type UserLoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthorizeUserLoginHandler(context echo.Context) error {
	if context.Request().Form == nil {
		context.Request().ParseForm()
	}
	u := UserLoginParam{}
	context.Bind(&u)
	if u.Username == "" || u.Password == "" {
		return context.JSON(status.Http400("用户名或密码不能为空", nil))
	}
	_, error := gws.GetSession(context.Response(), context.Request())
	if error != nil {
		return context.JSON(status.Http500("session错误", error))
	}
	context.Request().Form.Set("username", u.Username)
	context.Request().Form.Set("password", u.Password)
	context.Request().Form.Set("client_id", "auth")
	context.Request().Form.Set("client_secret", "auth")
	context.Request().Form.Set("grant_type", "password")
	context.Request().Form.Set("scope", "all")
	err := oauth2.Server.HandleTokenRequest(context.Response(), context.Request())
	if err != nil {
		return context.JSON(status.Http500("", err))
	}
	return nil
}
