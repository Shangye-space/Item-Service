package subcategory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

//Update - sub categories update
func Update(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	subCategoryID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if subCategoryID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	//Declare a new SubCategory struct
	var subCategory models.SubCategory

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&subCategory)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if subCategory.SubCategoryName != nil && len(*subCategory.SubCategoryName) > 0 {
		prop := fmt.Sprintf(`sub_category_name = "%v"`, *subCategory.SubCategoryName)
		setProp = append(setProp, prop)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	query := string(fmt.Sprintf("UPDATE sub_category SET %v WHERE category_id = %v;", strings.Join(setProp, ", "), subCategoryID))
	fmt.Println(query)
	db.Exec(query)

}
