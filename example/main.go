//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js

package main

import (
	"log"
	"net/http"

	_ "github.com/thetechnick/statik/example/statik"
	"github.com/thetechnick/statik/fs"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
