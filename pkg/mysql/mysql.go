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
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
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

func getColumns() error {
	for i, table := range Tables {
		// Parameter binding does not work??
		rows, err := Db.Query("SHOW COLUMNS FROM " + table.Name)
		if err != nil {
			return err
		}
		for rows.Next() {
			var col Column
			rows.Scan(
				&col.Field,
				&col.Type,
				&col.Null,
				&col.Key,
				&col.Default,
				&col.Extra,
			)
			Tables[i].Columns = append(Tables[i].Columns, col)
		}
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

	if err := getColumns(); err != nil {
		return err
	}

	// print debug
	dump()

	return nil
}
