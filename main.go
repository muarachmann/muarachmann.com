package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	// routing
	"github.com/gorilla/mux"
)

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

// Get default port
func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":3000"
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// variable := mux.Vars(r)
	// query := variable["q"]
	site_title := "<h1>Welcome to my website</h1>"
	if v := os.Getenv("APP_NAME"); v != "" {
		site_title = "<h1>Welcome to my" + v
	}
	fmt.Fprint(w, site_title)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, nil)
}

func main() {

	muxRouter := mux.NewRouter()

	fs := http.FileServer(http.Dir("public"))
	muxRouter.Handle("/public/", http.StripPrefix("/public/", fs))

	muxRouter.HandleFunc("/", indexHandler)
	muxRouter.HandleFunc("/about", aboutHandler)

	log.Fatal(http.ListenAndServe(getPort(), muxRouter))

}
