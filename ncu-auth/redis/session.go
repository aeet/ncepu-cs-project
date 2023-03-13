package redis

import (
	"fmt"
	"time"

	"github.com/auula/gws"
	"github.com/devcui/ncepu-cs-project/ncu-auth/config"
)

func SetupSession() {
	redis := config.Value.Redis
	session := config.Value.RedisSession
	option := &gws.RDSOption{
		Index:    uint8(session.Index),
		Prefix:   session.Prefix,
		Address:  fmt.Sprintf("%s:%d", redis.Host, redis.Port),
		Password: redis.Password,
		PoolSize: uint8(session.PoolSize),
	}
	option.CookieName = session.CookieName
	option.LifeTime = time.Duration(session.LifeTime) * time.Second
	option.HttpOnly = session.HttpOnly
	option.Secure = session.Secure
	option.Domain = session.Domain
	option.Path = session.Path
	gws.Open(option)
}
