package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Value config

type config struct {
	Auth         authOption         `yaml:"auth"`
	Server       serverOption       `yaml:"server"`
	Jwt          jWTOption          `yaml:"jwt"`
	Redis        redisOption        `yaml:"redis"`
	RedisSession redisSessionOption `yaml:"redis_session"`
	MySQLOption  mySQLOption        `yaml:"mysql"`
}

func Setup() error {
	path := flag.String("config", "config.yaml", "the absolute path of config.yaml")
	flag.Parse()

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &Value)
	if err != nil {
		return err
	}
	return nil
}
