package config

type serverOption struct {
	Path string `yaml:"path"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
