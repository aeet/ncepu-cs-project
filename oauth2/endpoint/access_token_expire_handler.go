package endpoint

import (
	"net/http"
	"time"

	"github.com/devcui/ncepu-cs-project/config"
	"github.com/labstack/gommon/log"
)

func AccessTokenExpireHandler(w http.ResponseWriter, r *http.Request) (exp time.Duration, err error) {
	log.Infof("AccessTokenExpireHandler")
	option := config.Value.Auth
	exp = time.Duration(option.AccessTokenExpire) * time.Second
	err = nil
	return
}
