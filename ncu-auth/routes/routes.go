package routes

import (
	"github.com/devcui/ncepu-cs-project/ncu-auth/routes/auth"
	"github.com/devcui/ncepu-cs-project/ncu-auth/routes/callback"
	"github.com/labstack/echo"
)

func Routes(api *echo.Group) {
	auth.Routes(api)
	callback.Routes(api.Group("/callback"))
}
