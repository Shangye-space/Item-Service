package subcategory

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

//UpdateHandler sub categories
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(w)

	subCategoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//Declare a new SubCategory struct
	var subCategory models.SubCategory

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&subCategory)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if helpers.CheckString(subCategory.Name) == nil {
		prop := fmt.Sprintf(`name = "%v"`, *subCategory.Name)
		setProp = append(setProp, prop)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Update(subCategoryID, db, setProp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

//Update - sub categories update
func Update(subCategoryID int, db *sql.DB, setProp []string) {
	query := string(fmt.Sprintf("UPDATE sub_category SET %v WHERE id = %v;", strings.Join(setProp, ", "), subCategoryID))
	db.Exec(query)
}
