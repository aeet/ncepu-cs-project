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
	// certificate
	api.POST("/certificate", CertificateAddHandler)
	api.GET("/certificate", CertificateListHandler)
	api.DELETE("/certificate/:id", CertificateDeleteHandler)
	api.PUT("/certificate", CertificateUpdateHandler)
	// family
	api.POST("/family", FamilyAddHandler)
	api.GET("/family", FamilyListHandler)
	api.DELETE("/family/:id", FamilyDeleteHandler)
	api.PUT("/family", FamilyUpdateHandler)
	// practical
	api.POST("/practical", PracticalAddHandler)
	api.GET("/practical", PracticalListHandler)
	api.DELETE("/practical/:id", PracticalDeleteHandler)
	api.PUT("/practical", PracticalUpdateHandler)
	// student
	api.POST("", StudentAddHandler)
	api.GET("", StudentListHandler)
	api.DELETE(":id", StudentDeleteHandler)
	api.PUT("", StudentUpdateHandler)
}
