package config

type redisSessionOption struct {
	Debug      bool   `yaml:"debug"`
	Index      int    `yaml:"index"`
	Prefix     string `yaml:"prefix"`
	Domain     string `yaml:"domain"`
	Path       string `yaml:"path"`
	CookieName string `yaml:"cookie_name"`
	HttpOnly   bool   `yaml:"http_only"`
	Secure     bool   `yaml:"secure"`
	LifeTime   int    `yaml:"life_time"`
	PoolSize   int    `yaml:"pool_size"`
}
