package callback

import (
	"github.com/labstack/echo"
)

func Setup(api *echo.Group) {
	callbackAPI := api.Group("/callback")
	routes(callbackAPI)
}

func routes(api *echo.Group) {
	api.Any("/show", CallbackShowHandler)
}
