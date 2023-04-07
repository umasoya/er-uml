package connection

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/umasoya/er-uml/pkg/config"
)

type mysqlConn struct{}

var mysql mysqlConn

// get tables name
func (conn *mysqlConn) getTables() error {
	result, err := db.Query("SHOW TABLES")
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

// get columns detail
func (conn *mysqlConn) getColumns() error {
	for i, table := range Tables {
		// Parameter binding does not work??
		rows, err := db.Query("SHOW COLUMNS FROM " + table.Name)
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

// Dump thw schema in json
func (conn *mysqlConn) dump() {
	json, err := json.MarshalIndent(Tables, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}

func (conn *mysqlConn) Run() error {
	// connection open
	if err := open(&config.Conf); err != nil {
		return err
	}

	// connection close
	defer close()

	if err := mysql.getTables(); err != nil {
		return err
	}

	if err := mysql.getColumns(); err != nil {
		return err
	}

	// print debug
	mysql.dump()

	return nil
}
