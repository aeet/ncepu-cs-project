package student

import (
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

// 证书
func CertificateAddHandler(c echo.Context) error {
	param := domain.Certificate{}
	c.Bind(&param)
	e := service.CertificateAdd(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func CertificateListHandler(c echo.Context) error {
	data, e := service.CertificateQuery()
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", data))
}

func CertificateDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(status.Http500(err.Error(), nil))
	}
	err = service.CertificateDelete(intID)
	if err != nil {
		return c.JSON(status.Http500(err.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func CertificateUpdateHandler(c echo.Context) error {
	param := domain.Certificate{}
	c.Bind(&param)
	e := service.CertificateUpdate(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

// 家庭
func FamilyAddHandler(c echo.Context) error {
	param := domain.FamilyInfo{}
	c.Bind(&param)
	e := service.FamilyAdd(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func FamilyListHandler(c echo.Context) error {
	data, e := service.FamilyQuery()
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", data))
}

func FamilyDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(status.Http500(err.Error(), nil))
	}
	e := service.FamilyDelete(intID)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func FamilyUpdateHandler(c echo.Context) error {
	param := domain.FamilyInfo{}
	c.Bind(&param)
	e := service.FamilyUpdate(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

// 实践
func PracticalAddHandler(c echo.Context) error {
	param := domain.PracticalExperience{}
	c.Bind(&param)
	e := service.PracticalAdd(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func PracticalListHandler(c echo.Context) error {
	data, e := service.PracticalQuery()
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", data))
}

func PracticalDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(status.Http500(err.Error(), nil))
	}
	e := service.PracticalDelete(intID)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func PracticalUpdateHandler(c echo.Context) error {
	param := domain.PracticalExperience{}
	c.Bind(&param)
	e := service.PracticalUpdate(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

// 学生
func StudentAddHandler(c echo.Context) error {
	param := domain.Student{}
	c.Bind(&param)
	e := service.StudentAdd(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func StudentListHandler(c echo.Context) error {
	data, e := service.StudentQuery()
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", data))
}

func StudentDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(status.Http500(err.Error(), nil))
	}
	e := service.StudentDelete(intID)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}

func StudentUpdateHandler(c echo.Context) error {
	param := domain.Student{}
	c.Bind(&param)
	e := service.StudentUpdate(param)
	if e != nil {
		return c.JSON(status.Http500(e.Error(), nil))
	}
	return c.JSON(status.Http200("success", nil))
}
