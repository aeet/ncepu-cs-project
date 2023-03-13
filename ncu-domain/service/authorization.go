package service

import (
	"context"

	"ncu-domain/domain"
	"ncu-domain/domain/authorization"
	"ncu-domain/mysql"
)

func Authorizations() (data []*domain.Authorization, err error) {
	client, e := mysql.Open()
	if e != nil {
		err = e
		return
	}
	data, err = client.Authorization.Query().All(context.Background())
	client.Close()
	return
}

func AuthorizationByClientIDAndClientSecret(clientID string, clientSecret string) (data *domain.Authorization, err error) {
	client, e := mysql.Open()
	if e != nil {
		err = e
		return
	}
	data, err = client.Authorization.Query().Where(authorization.ClientID(clientID), authorization.ClientSecret(clientSecret)).First(context.Background())
	client.Close()
	return
}

func AuthorizationExists(clientID string) (exist bool, err error) {
	client, e := mysql.Open()
	if e != nil {
		err = e
		return
	}
	exist, err = client.Authorization.Query().Where(authorization.ClientID(clientID)).Exist(context.Background())
	client.Close()
	return
}

func AuthorizationByClientID(clientID string) (data *domain.Authorization, err error) {
	client, e := mysql.Open()
	if e != nil {
		err = e
		return
	}
	data, err = client.Authorization.Query().Where(authorization.ClientID(clientID)).First(context.Background())
	client.Close()
	return
}

func AuthorizationScopeByClientID(clientID string) (scope []string, err error) {
	data, e := AuthorizationByClientID(clientID)
	if e != nil {
		err = e
		return
	}
	return data.Scope, nil
}

func AuthorizationFilterScopeByClientID(clientID string, requestScope []string) (scope []string, err error) {
	dbScope, e := AuthorizationScopeByClientID(clientID)
	if e != nil {
		err = e
		return
	}
	scopes := []string{}
	if err != nil {
		return nil, err
	}
	for _, rc := range requestScope {
		for _, dc := range dbScope {
			if rc == dc {
				scopes = append(scopes, rc)
			}
		}
	}
	return scopes, nil
}
