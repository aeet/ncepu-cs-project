package endpoint

import (
	"net/http"

	"github.com/devcui/ncepu-cs-project/status"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/gommon/log"
)

func PreRedirectErrorHandler(w http.ResponseWriter, req *server.AuthorizeRequest, err error) error {
	log.Errorf("PreRedirectErrorHandler: %s", req.Request.RequestURI)
	status.WriteErrorResponse(w, err)
	return nil
}
