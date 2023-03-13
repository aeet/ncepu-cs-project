package config

type authOption struct {
	AccessTokenExpire  int `yaml:"access_token_expire"`
	RefreshTokenExpire int `yaml:"refresh_token_expire"`
}
