package callback

import (
	"github.com/labstack/echo"
)

func Routes(api *echo.Group) {
	api.Any("/show", CallbackShowHandler)
}
