package image

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

//ReturnImageLink - returns link to image
func ReturnImageLink(w http.ResponseWriter, r *http.Request) {

	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	imageInstance, err := GetByID(itemID, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := imageInstance[0].ItemID
	link := "http://localhost:3348/api/image/return/" + strconv.Itoa(*id)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(link)
}
