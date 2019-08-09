package action

import (
	"log"
	"net/http"
	"text/template"
)

func InputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("view/index.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}
}
