package role

import (
	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func RoleAddHandler(c echo.Context) error {
	var role domain.Role
	c.Bind(&role)
	e := service.RoleAdd(role)
	if e != nil {
		return c.JSON(status.Http500("添加角色失败", e.Error()))
	}
	return c.JSON(status.Http200("添加角色成功", nil))
}

func RoleDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(status.Http400("参数错误", nil))
	}
	err := service.RoleDelete(id)
	if err != nil {
		return c.JSON(status.Http500("删除角色失败", err.Error()))
	}
	return c.JSON(status.Http200("删除角色成功", nil))
}

func RoleUpdateHandler(c echo.Context) error {
	var role domain.Role
	c.Bind(&role)
	e := service.RoleUpdate(role)
	if e != nil {
		return c.JSON(status.Http500("更新角色失败", e.Error()))
	}
	return c.JSON(status.Http200("更新角色成功", nil))
}

func RoleQueryHandler(c echo.Context) error {
	roles, err := service.RoleQuery()
	if err != nil {
		return c.JSON(status.Http500("查询角色失败", err.Error()))
	}
	return c.JSON(status.Http200("查询角色成功", roles))
}
