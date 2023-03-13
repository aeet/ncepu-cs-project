package oauth2

import (
	"github.com/devcui/ncepu-cs-project/domain/service"
	"github.com/devcui/ncepu-cs-project/oauth2/endpoint"
	ori "github.com/devcui/ncepu-cs-project/redis"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

var Manager *manage.Manager
var TokenStore *ori.TokenStore
var ClientStore *store.ClientStore
var Server *server.Server

func SetupServer() error {
	// Manager
	Manager, err := NewManager()
	if err != nil {
		return err
	}
	// Manager TokenStore
	TokenStore = NewTokenStore()
	Manager.MustTokenStorage(TokenStore, nil)
	Manager.MapAccessGenerate(NewJWTAccessGenerate())
	// Manager ClientStore
	ClientStore = store.NewClientStore()
	authorizations, err := service.Authorizations()
	if err != nil {
		return err
	}
	MapAuthorizationToStore(ClientStore, authorizations)
	Manager.MapClientStorage(ClientStore)
	// Server
	Server = server.NewServer(server.NewConfig(), Manager)
	Server.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.PasswordCredentials, oauth2.ClientCredentials, oauth2.Refreshing, oauth2.Implicit)
	Server.SetAllowGetAccessRequest(true)
	// Server.SetAccessTokenExpHandler(endpoint.AccessTokenExpireHandler)
	Server.SetPasswordAuthorizationHandler(endpoint.PasswordAuthorizationHandler)
	Server.SetUserAuthorizationHandler(endpoint.UserAuthorizationHandler)
	Server.SetAuthorizeScopeHandler(endpoint.AuthorizeScopeHandler)
	Server.SetClientInfoHandler(endpoint.ClientInfoHandler)
	Server.SetPreRedirectErrorHandler(endpoint.PreRedirectErrorHandler)
	Server.SetInternalErrorHandler(endpoint.InternalErrorHandler)
	// Server.SetResponseErrorHandler(endpoint.ResponseErrorHandler)
	return nil
}
