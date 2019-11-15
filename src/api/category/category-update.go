package category

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

//Update categories
func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	categoryID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if categoryID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	//Declare a new Category struct.
	var category models.Category

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&category)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if category.CategoryName != nil && len(*category.CategoryName) > 0 {
		prop := fmt.Sprintf(`category_name = "%v"`, *category.CategoryName)
		setProp = append(setProp, prop)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	query := string(fmt.Sprintf("UPDATE category SET %v WHERE category_id = %v;", strings.Join(setProp, ", "), categoryID))
	fmt.Println(query)
	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
