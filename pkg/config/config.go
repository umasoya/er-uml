package config

type Config struct {
	Schema   string `env:"SCHEMA"`
	User     string `env:"USER"     envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Host     string `env:"HOST"     envDefault:"localhost"`
	Port     string `env:"PORT"     envDefault:"3306"`
}

var Conf Config
