package department

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func DepartmentAddHandler(c echo.Context) error {
	var department domain.Department
	c.Bind(&department)
	e := service.DepartmentAdd(department)
	if e != nil {
		return c.JSON(status.Http500("添加系部失败", e.Error()))
	}
	return c.JSON(status.Http200("添加系部成功", nil))
}

func DepartmentDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(status.Http400("参数错误", nil))
	}
	err := service.DepartmentDelete(id)
	if err != nil {
		return c.JSON(status.Http500("删除系部失败", err.Error()))
	}
	return c.JSON(status.Http200("删除系部成功", nil))
}

func DepartmentUpdateHandler(c echo.Context) error {
	var department domain.Department
	c.Bind(&department)
	e := service.DepartmentUpdate(department)
	if e != nil {
		return c.JSON(status.Http500("更新系部失败", e.Error()))
	}
	return c.JSON(status.Http200("更新系部成功", nil))
}

func DepartmentQueryHandler(c echo.Context) error {
	departmentes, err := service.DepartmentQuery()
	if err != nil {
		return c.JSON(status.Http500("查询系部失败", err.Error()))
	}
	return c.JSON(status.Http200("查询系部成功", departmentes))
}
