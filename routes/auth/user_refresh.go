package auth

import (
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

type UserRefreshParam struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" query:"refresh_token"`
}

func AuthorizeUserRefreshHandler(context echo.Context) error {
	if context.Request().Form == nil {
		context.Request().ParseForm()
	}
	var token *UserRefreshParam = &UserRefreshParam{}
	context.Bind(token)
	if token == nil {
		return context.JSON(status.Http400("token不能为空", nil))
	}
	context.Request().Form.Set("client_id", "auth")
	context.Request().Form.Set("client_secret", "auth")
	context.Request().Form.Set("grant_type", "refresh_token")
	context.Request().Form.Set("scope", "all")
	context.Request().Form.Set("refresh_token", token.RefreshToken)
	err := oauth2.Server.HandleTokenRequest(context.Response(), context.Request())
	if err != nil {
		return context.JSON(status.Http500("", err))
	}
	return nil
}
