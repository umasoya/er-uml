package config

import (
	"fmt"
)

type Config struct {
	Driver   string `env:"DRIVER"   envDefault:"mysql"`
	User     string `env:"USER"     envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Host     string `env:"HOST"     envDefault:"localhost"`
	Port     string `env:"PORT"     envDefault:"3306"`
	Db       string `env:"DB"`
}

var Conf Config

func (c Config) Dsn() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?interpolateParams=true", c.User, c.Password, c.Host, c.Port, c.Db)
}
