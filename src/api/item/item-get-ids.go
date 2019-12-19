package item

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

//GetItemsByIDs - get item info by id
func GetItemsByIDs(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	params := mux.Vars(r)
	k := strings.Replace(params["ids"], " ", "", -1)
	s := strings.Split(k, ",")

	var items []models.Item
	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for _, element := range s {
		itemID, err := strconv.Atoi(element)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = helpers.CheckID(&itemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		item := GetByID(itemID, db)
		items = append(items, item[0])
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
