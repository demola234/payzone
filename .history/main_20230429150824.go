package main

import "database/sql"

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/payzone?sslmode=disable"
)


func main() {
	conn, err sql.Open(dbDriver, dbSource)

}