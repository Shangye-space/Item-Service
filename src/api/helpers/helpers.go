package helpers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
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

//EnableCors - enables Cross-Origin Resource Sharing
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
   }

//CheckIDWithRequest - checks whether id is valid with request
func CheckIDWithRequest(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		err = errors.New("Error occured while converting ID to int")
	} else if id <= 0 {
		err = errors.New("ID can't be negative or 0")
	}
	return id, err
}


//CheckID - checks whether id is valid
func CheckID(id *int) error {
	var err error
	if id == nil {
		err = errors.New("There is a problem with id")
	} else if *id <= 0 {
		err = errors.New("ID can't be negative or 0")
	}
	return err
}

//CheckBool - checks boolean and returns 0/1
func CheckBool(boolToCheck *bool) (int, error) {
	var num int
	var err error
	if boolToCheck == nil {
		err = errors.New("There is a problem with boolean")
	} else if *boolToCheck == true {
		num = 1
	} else {
		num = 0
	}
	return num, err
}

//CheckString - checks whether text is valid
func CheckString(text *string) error {
	var err error
	if text == nil || len(*text) == 0 {
		err = errors.New("There was a problem with text")
	}
	return err
}

//CheckNumber - checks whether number is valid
func CheckNumber(number *float32) error {
	var err error
	if number == nil || *number <= 0 {
		err = errors.New("There was a problem with number")
	}
	return err
}

//ScanItems - scans result for item
func ScanItems(result *sql.Rows) []models.Item {
	var item models.Item
	var items []models.Item
	for result.Next() {
		err := result.Scan(&item.ID, &item.Name, &item.Price, &item.SubCategoryID, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}
	return items
}

//ScanImage - scans result for image 
func ScanImage(result *sql.Rows) []models.Image{
	var image models.Image
	var images []models.Image
	for result.Next(){
		err := result.Scan(&image.ItemID, &image.Path)
		if err != nil {
			panic(err.Error())
		}
		images = append(images, image)
	}
	return images
}