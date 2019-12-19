package image

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Upload - uploads images
func Upload(r *http.Request) string {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)

	tempFile, err := ioutil.TempFile("./src/assets", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	path := tempFile.Name()

	fmt.Println("Image uploaded")
	return path
}
