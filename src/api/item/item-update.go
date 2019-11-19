package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

	if item.Name != nil && len(*item.Name) > 0 {
		prop := fmt.Sprintf(`name = "%v"`, *item.Name)
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

	// saves time of last update operation in db
	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	query := string(fmt.Sprintf("UPDATE item SET %v WHERE id = %v;", strings.Join(setProp, ", "), strconv.Itoa(itemID)))
	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
