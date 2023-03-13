package auth

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/auula/gws"
	"github.com/devcui/ncepu-cs-project/ncu-auth/config"
	"github.com/devcui/ncepu-cs-project/ncu-auth/status"
	"github.com/devcui/ncepu-cs-project/ncu-domain/domain"
	"github.com/devcui/ncepu-cs-project/ncu-domain/service"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func AuthorizeLoginHandler(context echo.Context) error {
	log.Infof("AuthorizeLoginHandler: %s", context.Request().URL.Path)
	session, e := gws.GetSession(context.Response(), context.Request())
	if e != nil {
		return context.JSON(status.Http500(e.Error(), nil))
	}

	// validate session
	authForm, ok := session.Values["AuthForm"]
	if !ok || authForm == nil {
		return context.JSON(status.Http500("session is invalid", nil))
	}

	// handle login request
	if context.Request().Method == "POST" {
		loginType := context.FormValue("type")

		// handle password login
		if loginType == "password" {
			account := context.FormValue("username")
			passwd := context.FormValue("password")
			user, err := service.User(account, passwd)
			var TplData map[string]interface{} = map[string]interface{}{}
			if err != nil {
				TplData["Error"] = err
				return context.Render(500, "login.html", TplData)
			}
			if user == nil {
				TplData["Error"] = errors.New("invalid username or password")
				return context.Render(500, "login.html", TplData)
			}
			// set session
			session.Values["LoggedInUserID"] = strconv.Itoa(user.ID)
			session.Sync()
			context.Redirect(302, config.Value.Server.Path)
		}

	}
	// get form from session
	formValues := url.Values{}
	authFormBytes, _ := json.Marshal(authForm)
	json.Unmarshal(authFormBytes, &formValues)
	// find authorization and render login page
	authorization, err := service.AuthorizationByClientID(formValues.Get("client_id"))
	if err != nil {
		return context.JSON(status.Http500(err.Error(), nil))
	}
	return context.Render(200, "login.html", templateDataFromAuthorizationAndForm(authorization, formValues))
}

func templateDataFromAuthorizationAndForm(authorization *domain.Authorization, form url.Values) map[string]interface{} {
	return map[string]interface{}{
		"client_id":   authorization.ClientID,
		"client_name": authorization.ClientName,
		"scope":       strings.Split(form.Get("scope"), ","),
	}
}
