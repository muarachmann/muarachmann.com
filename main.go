package main

import (
	"fmt"
	"net/http"
	"os"

	// router
	"github.com/gorilla/mux"
)

func handlerFucn(w http.ResponseWriter, r *http.Request) {
	// variable := mux.Vars(r)
	// query := variable["q"]
	site_title := "<h1>Welcome to my website</h1>"
	if v := os.Getenv("APP_NAME"); v != "" {
		site_title = "<h1>Welcome to my" + v
	}
	fmt.Fprint(w, site_title)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handlerFucn)
	http.ListenAndServe(":8080", router)

}
