package auth

import "github.com/labstack/echo"

func Routes(api *echo.Group) {
	api.Any("", AuthorizeHandler)
	api.Any("/login", AuthorizeLoginHandler)
	api.Any("/logout", AuthorizeLogoutHandler)
	api.Any("/token", AuthorizeTokenHandler)
	api.Any("/verify", AuthorizeVerifyHandler)
}
