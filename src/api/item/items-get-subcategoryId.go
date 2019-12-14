package item

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/Shangye-space/Item-Service/src/models"
// 	"github.com/gorilla/mux"

// 	database "github.com/Shangye-space/Item-Service/src/db"
// )

// // GetBySubCategoryID - Gets Item by SubCategory ID
// func GetBySubCategoryID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	subcategoryID, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	} else if subcategoryID == 0 {
// 		http.Error(w, "can't be 0", http.StatusBadRequest)
// 	}

// 	db, err := database.CreateDatabase()
// 	if err != nil {
// 		log.Fatal("Connection to DB has failed.")
// 	}

// 	fmt.Println(subcategoryID)
// 	query := string(fmt.Sprintf("SELECT * FROM item WHERE sub_category_id = %v", strconv.Itoa(subcategoryID)))
// 	result, err := db.Query(query)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	var item models.Item
// 	var items []models.Item

// 	for result.Next() {
// 		err := result.Scan(&item.ID, &item.Name, &item.SubCategoryID, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		items = append(items, item)
// 	}

// 	defer result.Close()

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(items)
// }
