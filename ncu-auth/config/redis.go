package config

type redisOption struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
	Expire   int    `yaml:"expire"`
}
