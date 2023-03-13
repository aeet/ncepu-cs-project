package student

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	studentAPI := api.Group("/student")
	studentAPI.Use(middleware.ValidateTokenMiddleware)
	routes(studentAPI)
}

func routes(api *echo.Group) {
	api.GET("/list", StudentHandler)
}
