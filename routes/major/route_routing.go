package major

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	roleAPI := api.Group("/major", middleware.ValidateTokenMiddleware)
	routes(roleAPI)
}

func routes(api *echo.Group) {
	api.POST("", MajorAddHandler)
	api.GET("", MajorQueryHandler)
	api.DELETE("/:id", MajorDeleteHandler)
	api.PUT("", MajorUpdateHandler)
}
