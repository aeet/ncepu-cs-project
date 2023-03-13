package auth

import (
	"encoding/json"
	"net/url"

	"github.com/auula/gws"
	"github.com/devcui/ncepu-cs-project/ncu-auth/oauth2"
	"github.com/devcui/ncepu-cs-project/ncu-auth/status"
	"github.com/labstack/echo"
)

func AuthorizeHandler(context echo.Context) error {
	context.Logger().Info("AuthorizeHandler: " + context.Request().URL.Path)
	session, e := gws.GetSession(context.Response(), context.Request())
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))

	}
	// 重定向到此Handler时: 尝试从session中获取AuthForm,如果可以获取则解析到form中
	var form url.Values
	if authForm, ok := session.Values["AuthForm"]; ok && authForm != nil {
		contextForm, e := context.FormParams()
		if e != nil {
			return context.JSON(status.Http500(e.Error(), nil))
		}
		if !contextForm.Has("client_id") {
			formBytes, _ := json.Marshal(session.Values["AuthForm"])
			json.Unmarshal(formBytes, &form)
		}
	}
	// 清空session中的AuthForm
	delete(session.Values, "AuthForm")
	err := session.Sync()
	if err != nil {
		return context.JSON(status.Http500(err.Error(), nil))
	}
	// 重定向Handler: 解析好的form放入context中
	context.Request().Form = form
	context.Request().PostForm = form
	err = oauth2.Server.HandleAuthorizeRequest(context.Response(), context.Request())
	if err != nil {
		return context.JSON(status.Http400(err.Error(), nil))
	}
	return nil
}
