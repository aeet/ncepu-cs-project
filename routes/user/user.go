package user

import (
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
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

func UserAdd(context echo.Context) error {
	params := domain.User{}
	context.Bind(&params)
	e := service.UserAdd(params)
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))
	}
	return context.JSON(status.Http200("success", nil))
}

func UserUpdate(context echo.Context) error {
	params := domain.User{}
	context.Bind(&params)
	e := service.UserUpdate(params)
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))
	}
	return context.JSON(status.Http200("success", nil))
}

func UserDelete(context echo.Context) error {
	id := context.Param("id")
	intID, e := strconv.Atoi(id)
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))
	}
	e = service.UserDelete(intID)
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))
	}
	return context.JSON(status.Http200("success", nil))
}

func UserQuery(context echo.Context) error {
	users, e := service.UserQuery()
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))
	}
	return context.JSON(status.Http200("success", users))
}
