package role

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	roleAPI := api.Group("/role", middleware.ValidateTokenMiddleware)
	routes(roleAPI)
}

func routes(api *echo.Group) {
	api.POST("", RoleAddHandler)
	api.GET("", RoleQueryHandler)
	api.DELETE("/:id", RoleDeleteHandler)
	api.PUT("", RoleUpdateHandler)
}
