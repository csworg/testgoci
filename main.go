package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type PageData struct {
	name string
}

//CSW Simple HTTP Server to test GoSec
func main() {

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	websitename := "cw-website"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := PageData {
			name: websitename,
		}

		if r.URL.Path[1:] != "" {
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		} else {
			tmpl.Execute(w, data)
			//fmt.Fprintf(w, "Hi there, you have reached Chris simple Go web server")
		}
	})

	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}