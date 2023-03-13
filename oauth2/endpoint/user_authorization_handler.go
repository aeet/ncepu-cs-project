package endpoint

import (
	"fmt"
	"net/http"

	"github.com/auula/gws"
	"github.com/devcui/ncepu-cs-project/config"
	"github.com/labstack/gommon/log"
)

func UserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (id string, err error) {
	log.Infof("UserAuthorizationHandler: %s", r.URL.Path)
	session, e := gws.GetSession(w, r)
	if e != nil {
		err = e
		return
	}
	// if user
	userID, ok := session.Values["LoggedInUserID"]
	if !ok || userID == nil {
		if r.Form == nil {
			r.ParseForm()
		}
		session.Values["AuthForm"] = r.Form
		e = session.Sync()
		if e != nil {
			err = e
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s/auth/login", config.Value.Server.Path))
		w.WriteHeader(http.StatusFound)
		return
	}
	id = userID.(string)
	return
}
