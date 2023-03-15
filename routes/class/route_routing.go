package class

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	roleAPI := api.Group("/class", middleware.ValidateTokenMiddleware)
	routes(roleAPI)
}

func routes(api *echo.Group) {
	api.POST("", ClassAddHandler)
	api.GET("", ClassQueryHandler)
	api.DELETE("/:id", ClassDeleteHandler)
	api.PUT("", ClassUpdateHandler)
}
