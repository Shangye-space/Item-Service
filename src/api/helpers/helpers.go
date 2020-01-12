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
func EnableCors(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
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

//CheckNumber - checks whether number (float) is valid
func CheckNumber(number *float32) error {
	var err error
	if number == nil || *number <= 0 {
		err = errors.New("There was a problem with number")
	}
	return err
}

//CheckNumberInt - checks whether number (int) is valid
func CheckNumberInt(number *int) error {
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
func ScanImage(result *sql.Rows) []models.Image {
	var image models.Image
	var images []models.Image
	for result.Next() {
		err := result.Scan(&image.ItemID, &image.Path)
		if err != nil {
			panic(err.Error())
		}
		images = append(images, image)
	}
	return images
}

//ScanCategories - scans result for categories
func ScanCategories(result *sql.Rows) []models.Category {
	var category models.Category
	var categories []models.Category

	for result.Next() {
		err := result.Scan(&category.ID, &category.Name, &category.AddedTime, &category.LastUpdated, &category.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}

	return categories
}

//ScanSubCategories - scans result for sub categories
func ScanSubCategories(result *sql.Rows) []models.SubCategory {
	var subCategory models.SubCategory
	var subCategories []models.SubCategory

	for result.Next() {
		err := result.Scan(&subCategory.ID, &subCategory.Name, &subCategory.CategoryID, &subCategory.AddedTime, &subCategory.LastUpdated, &subCategory.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		subCategories = append(subCategories, subCategory)
	}

	return subCategories

}

//ScanItemInfo - scans result for item info
func ScanItemInfo(result *sql.Rows) []models.ItemInfo {
	var itemInfo models.ItemInfo
	var itemInfos []models.ItemInfo

	for result.Next() {
		err := result.Scan(&itemInfo.ItemID,
			&itemInfo.Quantity,
			&itemInfo.Description,
			&itemInfo.Discount,
			&itemInfo.Size,
			&itemInfo.Color,
			&itemInfo.Manufacturer,
			&itemInfo.ItemCode,
			&itemInfo.Material,
			&itemInfo.AddedTime,
			&itemInfo.LastUpdated,
			&itemInfo.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		itemInfos = append(itemInfos, itemInfo)
	}

	return itemInfos
}
