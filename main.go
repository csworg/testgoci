package main

import (
	"fmt"
	"net/http"
	"os"
)

//CSW Simple HTTP Server to test GoSec
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})

	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}