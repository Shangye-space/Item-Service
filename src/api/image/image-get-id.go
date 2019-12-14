package image

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetByID - Gets images by itemID
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	fmt.Println(itemID)
	query := string(fmt.Sprintf("SELECT * FROM image WHERE item_id = %v", strconv.Itoa(itemID)))
	result, err := db.Query(query)

	fmt.Println(result)

	if err != nil {
		panic(err.Error())
	}

	var image models.Image
	var images []models.Image

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
