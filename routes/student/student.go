package student

import (
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func StudentHandler(context echo.Context) error {
	return context.JSON(status.Http200("hahahah", context.QueryParams()))
}
