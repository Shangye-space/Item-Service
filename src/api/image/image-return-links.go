package image

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/gorilla/mux"
)

//ReturnImageLinks - returns links to images for an array of itemIDs
func ReturnImageLinks(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(w)

	params := mux.Vars(r)
	k := strings.Replace(params["ids"], " ", "", -1)
	s := strings.Split(k, ",")
	var imageLinks []string
	for _, element := range s {
		itemID, err := strconv.Atoi(element)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = helpers.CheckID(&itemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		path := "http://localhost:3348/api/image/return/" + strconv.Itoa(itemID)
		imageLinks = append(imageLinks, path)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(imageLinks)
}
