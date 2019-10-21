package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	serverName := "192.168.45.100:3306"
	user := "root"
	password := "root"
	dbName := "db"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, serverName, dbName)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return database, nil
}
