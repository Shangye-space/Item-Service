package image

import (
	"io/ioutil"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

//ReturnImage - returns image
func ReturnImage(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(w)

	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic("error with id")
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	imageInstance, err := GetByID(itemID, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	path := "./" + *imageInstance[0].Path
	img := ReadImageFromFile(path)

	defer db.Close()
	w.Header().Set("Content-type", "image/png")
	w.Write(img)
}

//ReadImageFromFile - reads image from file
func ReadImageFromFile(path string) []byte {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("There is a problem with reading file")
	}

	return file
}
