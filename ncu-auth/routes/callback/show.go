package callback

import (
	"github.com/devcui/ncepu-cs-project/ncu-auth/status"
	"github.com/labstack/echo"
)

func CallbackShowHandler(context echo.Context) error {
	return context.JSON(status.Http200("generate code succeed", context.QueryParams()))
}
