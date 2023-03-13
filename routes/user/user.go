package user

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func UserHandler(context echo.Context) error {
	user := context.Get("user").(*domain.User)
	if user != nil {
		return context.JSON(status.Http200("success", user))
	}
	return context.JSON(status.Http500("user was null!", nil))
}
