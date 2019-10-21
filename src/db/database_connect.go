package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	user := "root"
	password := "root"
	dbName := "db"

	connectionString := fmt.Sprintf("%s:%s@tcp(docker.for.mac.localhost:3308)/%s?parseTime=true", user, password, dbName)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return database, nil
}
