package connection

import "fmt"

type postgresConn struct{}

var postgres postgresConn

func (conn *postgresConn) getTables() error {
	return nil
}

func (conn *postgresConn) Run() error {
	conn.getTables()

	fmt.Println("postgres stub")

	return nil
}
