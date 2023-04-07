package connection

import (
	"database/sql"
	"errors"

	"github.com/umasoya/er-uml/pkg/config"
)

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

var (
	Tables []Table
	db     *sql.DB
)

// open db connection
func open(conf *config.Config) error {
	var err error = nil
	db, err = sql.Open(conf.Driver, conf.Dsn())
	return err
}

// close db connection
func close() error {
	return db.Close()
}

func Execute(conf *config.Config) error {
	if err := open(conf); err != nil {
		return err
	}
	defer close()

	switch conf.Driver {
	case "mysql":
		mysql.Run()
	case "postgres":
		postgres.Run()
	default:
		return errors.New("undefined database driver: " + conf.Driver)
	}

	return nil
}
