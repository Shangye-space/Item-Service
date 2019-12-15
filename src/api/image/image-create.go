package image

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Create - saves an image of item
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)
	
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)

	tempFile, err := ioutil.TempFile("./src/assets", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tempFile.Name())
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)


	
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
