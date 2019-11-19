package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// CreateDatabase - creates db connection
func CreateDatabase() (*sql.DB, error) {
	host := "host.docker.internal"
	user := "root"
	password := "root"
	dbName := "db"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3308)/%s?parseTime=true", user, password, host, dbName)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return database, nil
}
