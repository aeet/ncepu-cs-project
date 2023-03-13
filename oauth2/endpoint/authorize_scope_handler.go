package endpoint

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/gommon/log"
)

func AuthorizeScopeHandler(w http.ResponseWriter, r *http.Request) (scope string, err error) {
	log.Infof("AuthorizeScopeHandler: %s", r.URL.Path)
	if r.Form == nil {
		r.ParseForm()
	}
	formScope := r.Form.Get("scope")
	formClientID := r.Form.Get("client_id")
	scopes, e := service.AuthorizationFilterScopeByClientID(formClientID, strings.Split(formScope, ","))
	if err != nil {
		err = e
		return
	}
	if len(scopes) <= 0 {
		err = errors.New(fmt.Sprintf("invalid_scope: %s", formScope))
		return
	} else {
		scope = strings.Join(scopes, ",")
		return
	}
}
