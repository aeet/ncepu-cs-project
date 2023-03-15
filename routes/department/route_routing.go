package department

import (
	"github.com/devcui/ncepu-cs-project/middleware"
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	roleAPI := api.Group("/department", middleware.ValidateTokenMiddleware)
	routes(roleAPI)
}

func routes(api *echo.Group) {
	api.POST("", DepartmentAddHandler)
	api.GET("", DepartmentQueryHandler)
	api.DELETE("/:id", DepartmentDeleteHandler)
	api.PUT("", DepartmentUpdateHandler)
}
