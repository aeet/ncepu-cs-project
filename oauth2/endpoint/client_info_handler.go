package endpoint

import (
	"net/http"

	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/gommon/log"
)

func ClientInfoHandler(r *http.Request) (clientID, clientSecret string, err error) {
	log.Infof("ClientInfoHandler: %s", r.URL.Path)
	if r.Form == nil {
		r.ParseForm()
	}
	clientID = r.Form.Get("client_id")
	clientSecret = r.Form.Get("client_secret")
	if clientID == "" || clientSecret == "" {
		err = errors.ErrInvalidClient
		return
	}
	client, e := service.AuthorizationByClientIDAndClientSecret(clientID, clientSecret)
	if e != nil {
		err = e
		return
	}
	if client == nil {
		err = errors.ErrInvalidClient
		return
	}
	clientID = client.ClientID
	clientSecret = client.ClientSecret
	err = nil
	return
}
