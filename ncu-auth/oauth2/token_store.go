package oauth2

import (
	"fmt"

	"github.com/devcui/ncepu-cs-project/ncu-auth/config"
	ori "github.com/devcui/ncepu-cs-project/ncu-auth/redis"
	"github.com/go-redis/redis/v8"
)

func NewTokenStore() *ori.TokenStore {
	option := config.Value.Redis
	return ori.NewRedisStore(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", option.Host, option.Port),
		DB:       option.DB,
		Password: option.Password,
	})
}
