package category

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

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

	if category.CategoryName == nil || len(*category.CategoryName) <= 0 {
		http.Error(w, "CategoryName is wrong", http.StatusBadRequest)
	}

	query := fmt.Sprintf(`
	INSERT INTO category(category_name)
	VALUES(%v);`, *category.CategoryName)

	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}