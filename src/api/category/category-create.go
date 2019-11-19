package category

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

//Create - creates categories in db
func Create(w http.ResponseWriter, r *http.Request) {
	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var category models.Category

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&category)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}

	if category.Name == nil || len(*category.Name) <= 0 {
		http.Error(w, "Name is wrong", http.StatusBadRequest)
	}

	query := fmt.Sprintf(`
	INSERT INTO category(name)
	VALUES("%v");`, *category.Name)

	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
