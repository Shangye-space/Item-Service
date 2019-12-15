package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/gorilla/mux"
)

//CreateDatabase - creates db connection
func CreateDatabase() (*sql.DB, error) {
	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}
	return db, err
}

//CheckID - check whether id is valid
func CheckID(r *http.Request) (int, error) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Error occured while converting itemID to int")
	} else if itemID <= 0 {
		fmt.Println("can't be 0")
		err = errors.New("ItemID can't be negative or 0")
	}
	return itemID, err
}
