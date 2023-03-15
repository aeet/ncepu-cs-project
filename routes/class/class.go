package class

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func ClassAddHandler(c echo.Context) error {
	var class domain.Class
	c.Bind(&class)
	e := service.ClassAdd(class)
	if e != nil {
		return c.JSON(status.Http500("添加班级失败", e.Error()))
	}
	return c.JSON(status.Http200("添加班级成功", nil))
}

func ClassDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(status.Http400("参数错误", nil))
	}
	err := service.ClassDelete(id)
	if err != nil {
		return c.JSON(status.Http500("删除班级失败", err.Error()))
	}
	return c.JSON(status.Http200("删除班级成功", nil))
}

func ClassUpdateHandler(c echo.Context) error {
	var class domain.Class
	c.Bind(&class)
	e := service.ClassUpdate(class)
	if e != nil {
		return c.JSON(status.Http500("更新班级失败", e.Error()))
	}
	return c.JSON(status.Http200("更新班级成功", nil))
}

func ClassQueryHandler(c echo.Context) error {
	classes, err := service.ClassQuery()
	if err != nil {
		return c.JSON(status.Http500("查询班级失败", err.Error()))
	}
	return c.JSON(status.Http200("查询班级成功", classes))
}
