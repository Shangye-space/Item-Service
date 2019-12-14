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

// GetByID - GetByID Item Info
func GetByID(w http.ResponseWriter, r *http.Request) {

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

	query := string(fmt.Sprintf("SELECT * FROM item_info WHERE item_id = %v", strconv.Itoa(itemID)))
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

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

	defer result.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemInfos)
}
