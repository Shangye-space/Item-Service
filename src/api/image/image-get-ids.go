package image

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetByIDs - Gets images by itemIDs
func GetByIDs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	k := strings.Replace(params["ids"], " ", "", -1)
	s := strings.Split(k, ",")
	var itemIDs []string
	for _, element := range s {
		itemID, err := strconv.Atoi(element)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if itemID <= 0 {
			http.Error(w, "can't be 0 or negative", http.StatusBadRequest)
		}

		itemIDs = append(itemIDs, element)
	}
	listOfIDs := strings.Join(itemIDs, ", ")

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	fmt.Println(itemIDs)
	query := string(fmt.Sprintf("SELECT * FROM image WHERE item_id IN (%v)", listOfIDs))
	result, err := db.Query(query)

	fmt.Println(result)

	if err != nil {
		panic(err.Error())
	}

	var image models.Item
	var images []models.Item

	for result.Next() {
		err := result.Scan()
		if err != nil {
			panic(err.Error())
		}
		images = append(images, image)
	}

	defer result.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
