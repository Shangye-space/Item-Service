package subcategory

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

//Create - creates sub categories in db
func Create(w http.ResponseWriter, r *http.Request) {
	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var subCategory models.SubCategory

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&subCategory)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}

	if subCategory.SubCategoryName == nil || len(*subCategory.SubCategoryName) <= 0 {
		http.Error(w, "SubCategoryName is wrong", http.StatusBadRequest)
	}

	if subCategory.CategoryID == nil || *subCategory.CategoryID <= 0 {
		http.Error(w, "CategoryID is wrong", http.StatusBadRequest)
	}

	query := fmt.Sprintf(`INSERT INTO sub_category(sub_category_name, category_id)
	VALUES("%v", %v);`, *subCategory.SubCategoryName, *subCategory.CategoryID)

	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
