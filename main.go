package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(100 << 20)

	request := r.URL

	fmt.Printf("%v", request)

	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(handler.Filename)
	fmt.Println(handler.Size)
	fmt.Println(handler.Header)

	tempfile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempfile.Close()

	filebytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempfile.Write(filebytes)

	fmt.Println("Upload complete")
}

func main() {

	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
