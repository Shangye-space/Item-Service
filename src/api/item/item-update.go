package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

// Update - updates post
func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	// Declare a new Item struct.
	var item models.Item

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&item)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if item.ItemName != nil && len(*item.ItemName) > 0 {
		prop := fmt.Sprintf(`item_name = "%v"`, *item.ItemName)
		setProp = append(setProp, prop)
	}

	if item.InSale != nil {
		var prop string
		if *item.InSale == false {
			prop = fmt.Sprintf(`in_sale = %v`, 0)
		} else {
			prop = fmt.Sprintf(`in_sale = %v`, 1)
		}
		setProp = append(setProp, prop)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	query := string(fmt.Sprintf("UPDATE item SET %v WHERE item_id = %v;", strings.Join(setProp, ", "), strconv.Itoa(itemID)))
	fmt.Println(query)
	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
