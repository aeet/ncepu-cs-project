package routes

import (
	"github.com/devcui/ncepu-cs-project/routes/auth"
	"github.com/devcui/ncepu-cs-project/routes/callback"
	"github.com/devcui/ncepu-cs-project/routes/role"
	"github.com/devcui/ncepu-cs-project/routes/student"
	"github.com/devcui/ncepu-cs-project/routes/user"
	"github.com/labstack/echo"
)

func Routes(api *echo.Group) {
	auth.Setup(api)
	callback.Setup(api)
	student.Setup(api)
	user.Setup(api)
	role.Setup(api)
}
