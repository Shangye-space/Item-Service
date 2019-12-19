package image

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

//CreateHandler - handles saving an image of item
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(itemID)
	
	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	path := Upload(r)

	Create(itemID, path, db)

}

//Create - saves an image of item
func Create(itemID int , path string, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO image(item_id, path) VALUES(%v, "%v");`, itemID, path)
	db.Exec(query)
}
