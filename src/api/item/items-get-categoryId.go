package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

// GetByCategoryID - Gets Item by Category ID
func GetByCategoryID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	categoryID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if categoryID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	query := string(fmt.Sprintf("http://localhost:3348/api/sub_categories/category/%v", strconv.Itoa(categoryID)))
	response, err := http.Get(query)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var subCategory []models.SubCategory

	json.Unmarshal(responseData, &subCategory)
	fmt.Printf("Subcategorys: %+v", subCategory)

	var ids []string
	for _, element := range subCategory {
		ids = append(ids, strconv.Itoa(*element.ID))
		fmt.Print(777)
	}

	decoder := json.NewDecoder(response.Body)
	err1 := decoder.Decode(&subCategory)
	if err1 != nil {
		panic(err)
	}

	query = string(fmt.Sprintf("SELECT * FROM item WHERE sub_category_id IN (%v)", strings.Join(ids, ", ")))
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var item models.Item
	var items []models.Item

	for result.Next() {
		err := result.Scan(&item.ID, &item.Name, &item.SubCategoryID, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	defer result.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
