package oauth2

import (
	"time"

	"github.com/devcui/ncepu-cs-project/ncu-auth/config"
	"github.com/go-oauth2/oauth2/v4/manage"
)

func NewManager() (*manage.Manager, error) {
	option := config.Value.Auth
	mgr := manage.NewDefaultManager()
	mgr.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp:    time.Duration(option.AccessTokenExpire) * time.Second,
		RefreshTokenExp:   time.Duration(option.RefreshTokenExpire+20) * time.Second,
		IsGenerateRefresh: true,
	})
	mgr.SetClientTokenCfg(&manage.Config{
		AccessTokenExp:    time.Duration(option.AccessTokenExpire) * time.Second,
		RefreshTokenExp:   time.Duration(option.RefreshTokenExpire+20) * time.Second,
		IsGenerateRefresh: true,
	})
	mgr.SetImplicitTokenCfg(&manage.Config{
		AccessTokenExp:    time.Duration(option.AccessTokenExpire) * time.Second,
		RefreshTokenExp:   time.Duration(option.RefreshTokenExpire+20) * time.Second,
		IsGenerateRefresh: true,
	})
	mgr.SetPasswordTokenCfg(&manage.Config{
		AccessTokenExp:    time.Duration(option.AccessTokenExpire) * time.Second,
		RefreshTokenExp:   time.Duration(option.RefreshTokenExpire+20) * time.Second,
		IsGenerateRefresh: true,
	})
	mgr.SetRefreshTokenCfg(&manage.RefreshingConfig{
		AccessTokenExp:    time.Duration(option.AccessTokenExpire) * time.Second,
		RefreshTokenExp:   time.Duration(option.RefreshTokenExpire+20) * time.Second,
		IsGenerateRefresh: true,
	})
	return mgr, nil

}
