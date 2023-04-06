package connection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/umasoya/er-uml/pkg/config"
)

// open db connection
func Open(conf *config.Config) (*sql.DB, error) {
	return sql.Open(conf.Driver, conf.Dsn())
}
