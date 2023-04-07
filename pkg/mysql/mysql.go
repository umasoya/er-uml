package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/umasoya/er-uml/pkg/config"
)

var Tables []Table
var Db *sql.DB

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Field string
	Type  string
	NULL  bool
	Extra []Extra
}

type Extra struct {
	AutoIncrement bool
}

// open db connection
func open(conf *config.Config) error {
	var err error = nil
	Db, err = sql.Open("mysql", conf.Dsn())
	return err
}

// close db connection
func close() error {
	return Db.Close()
}

// get tables name
func getTables() error {
	result, err := Db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	for result.Next() {
		var table Table
		result.Scan(&table.Name)
		Tables = append(Tables, table)
	}
	return nil
}

func dump() {
	json, err := json.MarshalIndent(Tables, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}

func Run() error {
	// connection open
	if err := open(&config.Conf); err != nil {
		return err
	}

	// connection close
	defer close()

	if err := getTables(); err != nil {
		return err
	}

	// print debug
	dump()

	return nil
}
