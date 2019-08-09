package action

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

//UploadIndex 上傳圖檔 index
func UploadIndex(w http.ResponseWriter, r *http.Request) {
	uploadTemplate := template.Must(template.ParseFiles("view/upload.html"))
	uploadTemplate.Execute(w, nil)
}

// UploadHandle 上傳圖檔 action
func UploadHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UploadHandle")
	file, head, _ := r.FormFile("file")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	w.Write(bytes)
	ioutil.WriteFile(head.Filename, bytes, os.ModeAppend)
}
