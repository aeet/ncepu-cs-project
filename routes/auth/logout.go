package auth

import (
	"net/url"

	"github.com/auula/gws"
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func AuthorizeLogoutHandler(context echo.Context) error {
	log.Infof("AuthorizeLogoutHandler: %s", context.Request().URL.Path)
	redirectURL := context.QueryParam("redirect_uri")
	if redirectURL == "" {
		return context.JSON(status.Http400("redirect_uri is required", nil))
	}
	if _, err := url.Parse(redirectURL); err != nil {
		return context.JSON(status.Http400("redirect_uri is invalid", nil))
	}
	session, error := gws.GetSession(context.Response(), context.Request())
	if error != nil {
		return context.JSON(status.Http500("get session error", nil))
	}
	delete(session.Values, "LoggedInUserID")
	if err := session.Sync(); err != nil {
		return context.JSON(status.Http500("session sync error", nil))
	}
	token := context.QueryParam("token")
	if token == "" {
		return context.JSON(status.Http400("token is required in request header like Authorization: Bearer xxxxxx", nil))
	}
	tokenInfo, err := oauth2.TokenStore.GetByAccess(context.Request().Context(), token)
	if err != nil {
		return context.JSON(status.Http500("remove access token error", nil))
	}
	oauth2.TokenStore.RemoveByAccess(context.Request().Context(), tokenInfo.GetAccess())
	oauth2.TokenStore.RemoveByCode(context.Request().Context(), tokenInfo.GetCode())
	oauth2.TokenStore.RemoveByRefresh(context.Request().Context(), tokenInfo.GetRefresh())
	context.Redirect(302, redirectURL)
	return nil
}
