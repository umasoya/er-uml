package config

import "fmt"

type Mysql struct {
	User         string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

func (conf Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DatabaseName)
}
