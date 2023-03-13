package config

type mySQLOption struct {
	Init     bool   `yaml:"init"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Schema   string `yaml:"schema"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
