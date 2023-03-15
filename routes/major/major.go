package major

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func MajorAddHandler(c echo.Context) error {
	var major domain.Major
	c.Bind(&major)
	e := service.MajorAdd(major)
	if e != nil {
		return c.JSON(status.Http500("添加专业失败", e.Error()))
	}
	return c.JSON(status.Http200("添加专业成功", nil))
}

func MajorDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(status.Http400("参数错误", nil))
	}
	err := service.MajorDelete(id)
	if err != nil {
		return c.JSON(status.Http500("删除专业失败", err.Error()))
	}
	return c.JSON(status.Http200("删除专业成功", nil))
}

func MajorUpdateHandler(c echo.Context) error {
	var major domain.Major
	c.Bind(&major)
	e := service.MajorUpdate(major)
	if e != nil {
		return c.JSON(status.Http500("更新专业失败", e.Error()))
	}
	return c.JSON(status.Http200("更新专业成功", nil))
}

func MajorQueryHandler(c echo.Context) error {
	majores, err := service.MajorQuery()
	if err != nil {
		return c.JSON(status.Http500("查询专业失败", err.Error()))
	}
	return c.JSON(status.Http200("查询专业成功", majores))
}
