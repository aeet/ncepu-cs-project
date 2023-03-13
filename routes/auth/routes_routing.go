package auth

import "github.com/labstack/echo"

func Setup(api *echo.Group) {
	authAPI := api.Group("/auth")
	routes(authAPI)
}

func routes(api *echo.Group) {
	api.Any("", AuthorizeHandler)
	api.Any("/login", AuthorizeLoginHandler)
	api.Any("/logout", AuthorizeLogoutHandler)
	api.Any("/token", AuthorizeTokenHandler)
	api.Any("/verify", AuthorizeVerifyHandler)
	api.POST("/user/login", AuthorizeUserLoginHandler)
	api.POST("/user/refresh", AuthorizeUserRefreshHandler)
}
