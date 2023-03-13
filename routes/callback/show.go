package callback

import (
	status "github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func CallbackShowHandler(context echo.Context) error {
	return context.JSON(status.Http200("generate code succeed", context.QueryParams()))
}
