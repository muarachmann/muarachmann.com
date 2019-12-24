package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// routing
	"github.com/gorilla/mux"
	"github.com/muarachmann/muarachmann.com/app/controllers"
)

const (
	STATIC_DIR = "/public/"
	PORT       = "3000"
)

// Get default port
func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":" + PORT
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

func main() {

	muxRouter := mux.NewRouter()

	muxRouter.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	muxRouter.HandleFunc("/", controllers.GetHomePage)
	muxRouter.HandleFunc("/about", controllers.GetAboutPage)
	muxRouter.HandleFunc("/contact", controllers.GetContactPage)

	log.Fatal(http.ListenAndServe(getPort(), muxRouter))

}
