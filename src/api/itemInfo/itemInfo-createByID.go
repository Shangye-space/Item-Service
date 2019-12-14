package iteminfo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

// CreateByID - creates item info in db by ID
func CreateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't b 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Conection to DB has failed")
	}

	var iteminfo models.ItemInfo
	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&iteminfo)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}

	if iteminfo.Quantity == nil || *iteminfo.Quantity <= 0 {
		http.Error(w, "Quantity is wrong", http.StatusBadRequest)
	}

	if iteminfo.Description == nil || len(*iteminfo.Description) <= 0 {
		http.Error(w, "Description is wrong", http.StatusBadRequest)
	}

	if iteminfo.Discount == nil || *iteminfo.Discount <= 0 {
		http.Error(w, "Discount is wrong", http.StatusBadRequest)
	}

	if iteminfo.Size == nil || len(*iteminfo.Size) <= 0 {
		http.Error(w, "Size is wrong", http.StatusBadRequest)
	}

	if iteminfo.Color == nil || len(*iteminfo.Color) <= 0 {
		http.Error(w, "Color is wrong", http.StatusBadRequest)
	}

	if iteminfo.Manufacturer == nil || len(*iteminfo.Manufacturer) <= 0 {
		http.Error(w, "Manufacturer is wrong", http.StatusBadRequest)
	}

	if iteminfo.ItemCode == nil || len(*iteminfo.ItemCode) <= 0 {
		http.Error(w, "ItemCode is wrong", http.StatusBadRequest)
	}

	if iteminfo.Material == nil || len(*iteminfo.Material) <= 0 {
		http.Error(w, "Material is wrong", http.StatusBadRequest)
	}

	query := fmt.Sprintf(`INSERT INTO item_info (item_id, quantity, description, discount, size, color, manufacturer, item_code, material) 
	VALUES(%v, %v, "%v", %v, "%v", "%v", "%v", "%v", "%v")`, itemID, *iteminfo.Quantity, *iteminfo.Description,
		*iteminfo.Discount, *iteminfo.Size, *iteminfo.Color, *iteminfo.Manufacturer, *iteminfo.ItemCode, *iteminfo.Material)
	fmt.Printf(query)
	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
