package middleware

import (
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := oauth2.Server.ValidationBearerToken(c.Request())
		if err != nil {
			return c.JSON(status.Http401("invalid token!", err))
		}
		return next(c)
	}
}
