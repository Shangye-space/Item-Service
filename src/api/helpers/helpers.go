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

//CheckID - checks whether id is valid
func CheckID(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Error occured while converting ID to int")
	} else if id <= 0 {
		fmt.Println("can't be 0")
		err = errors.New("ID can't be negative or 0")
	}
	return id, err
}

//CheckBool - checks
func CheckBool(boolToCheck bool) int{
	var num int
	if boolToCheck == true {
		num = 1
	} else {
		num = 0
	}
	return num
}

//CheckString - check text
func CheckString(text *string) (*string, error) {
	var err error
	if text == nil || len(*text) <= 0 {
		err = errors.New("There was a problem with text")
	}
	return text, err
}

//ChechNumber - check number
func CheckNumber(number *float32) (*float32, error) {
	var err error
	if number == nil || *number <= 0 {
		err = errors.New("There was a number with number")
	}
	return number, err
}
