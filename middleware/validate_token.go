package middleware

import (
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/oauth2"
	"github.com/devcui/ncepu-cs-project/status"
	"github.com/labstack/echo"
)

func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenInfo, err := oauth2.Server.ValidationBearerToken(c.Request())
		if err != nil {
			return c.JSON(status.Http401("invalid token!", err))
		}
		userID := tokenInfo.GetUserID()
		if userID == "" {
			return c.JSON(status.Http500("token hasn't userID!", err))
		}
		toIntUserID, err := strconv.Atoi(userID)
		if err != nil {
			return c.JSON(status.Http500("token userID is not int!", err))
		}
		user, err := service.UserByID(toIntUserID)
		if err != nil {
			return c.JSON(status.Http500("user was null!", err))
		}
		c.Set("user", user)
		return next(c)
	}
}
