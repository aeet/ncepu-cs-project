package oauth2

import (
	"github.com/go-oauth2/oauth2/v4/manage"
)

func NewManager() (*manage.Manager, error) {
	mgr := manage.NewDefaultManager()
	return mgr, nil

}
