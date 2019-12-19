package category

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	helpers "github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

//CreateHandler - handles creating category
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	var category models.Category

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Create(category, db)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//Create - creates categories in db
func Create(category models.Category, db *sql.DB) {
	query := fmt.Sprintf(`
	INSERT INTO category(name)
	VALUES("%v");`, *category.Name)

	db.Exec(query)
}
