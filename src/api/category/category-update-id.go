package category

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"strconv"

	helpers "github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

//UpdateHandler categories
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	categoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//Declare a new Category struct.
	var category models.Category

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&category)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if helpers.CheckString(category.Name) == nil {
		prop := fmt.Sprintf(`name = "%v"`, *category.Name)
		setProp = append(setProp, prop)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Update(categoryID, db, setProp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

//Update - updates category
func Update(categoryID int, db *sql.DB, setProp []string) {
	query := string(fmt.Sprintf("UPDATE category SET %v WHERE id = %v;", strings.Join(setProp, ", "), strconv.Itoa(categoryID)))
	db.Exec(query)
}
