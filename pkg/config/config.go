package config

type Config struct {
	Driver   string `env:"Driver"   envDefault:"mysql"`
	User     string `env:"USER"     envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Host     string `env:"HOST"     envDefault:"localhost"`
	Port     string `env:"PORT"     envDefault:"3306"`
}

var Conf Config
