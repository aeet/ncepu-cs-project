package oauth2

import (
	"github.com/devcui/ncepu-cs-project/ncu-domain/domain"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

func MapAuthorizationToStore(store *store.ClientStore, authorizations []*domain.Authorization) {
	for _, client := range authorizations {
		store.Set(client.ClientID, &models.Client{
			ID:     client.ClientID,
			Secret: client.ClientSecret,
			Domain: client.Domain,
		})
	}
}
