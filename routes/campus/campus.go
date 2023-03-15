package campus

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func CampusAddHandler(c echo.Context) error {
	var campus domain.Campus
	c.Bind(&campus)
	e := service.CampusAdd(campus)
	if e != nil {
		return c.JSON(status.Http500("添加校区失败", e.Error()))
	}
	return c.JSON(status.Http200("添加校区成功", nil))
}

func CampusDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(status.Http400("参数错误", nil))
	}
	err := service.CampusDelete(id)
	if err != nil {
		return c.JSON(status.Http500("删除校区失败", err.Error()))
	}
	return c.JSON(status.Http200("删除校区成功", nil))
}

func CampusUpdateHandler(c echo.Context) error {
	var campus domain.Campus
	c.Bind(&campus)
	e := service.CampusUpdate(campus)
	if e != nil {
		return c.JSON(status.Http500("更新校区失败", e.Error()))
	}
	return c.JSON(status.Http200("更新校区成功", nil))
}

func CampusQueryHandler(c echo.Context) error {
	campuses, err := service.CampusQuery()
	if err != nil {
		return c.JSON(status.Http500("查询校区失败", err.Error()))
	}
	return c.JSON(status.Http200("查询校区成功", campuses))
}
