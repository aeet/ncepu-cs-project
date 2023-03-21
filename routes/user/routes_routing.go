package user

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	userAPI := api.Group("/user")
	userAPI.Use(middleware.ValidateTokenMiddleware)
	routes(userAPI)
}

func routes(api *echo.Group) {
	api.GET("", UserHandler)
	api.GET("/query", UserQuery)
	api.POST("", UserAdd)
	api.PUT("", UserUpdate)
	api.DELETE("/:id", UserDelete)
}
