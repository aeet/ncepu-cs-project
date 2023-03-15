package campus

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	roleAPI := api.Group("/campus", middleware.ValidateTokenMiddleware)
	routes(roleAPI)
}

func routes(api *echo.Group) {
	api.POST("", CampusAddHandler)
	api.GET("", CampusQueryHandler)
	api.DELETE("/:id", CampusDeleteHandler)
	api.PUT("", CampusUpdateHandler)
}
