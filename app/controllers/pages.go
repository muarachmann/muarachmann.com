package controllers

import (
	"html/template"
	"net/http"

	"github.com/muarachmann/muarachmann.com/app/common"
)

// function to get the homepage
func GetHomePage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title   string
		Navlink string
	}

	p := Page{
		Title:   "Home",
		Navlink: "home",
	}

	templ, err := template.New("").ParseFiles("templates/index.html", "templates/base.html")
	err = templ.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}

// function to get the about us page
func GetAboutPage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title   string
		Navlink string
	}

	p := Page{
		Title:   "About",
		Navlink: "about",
	}

	templ, err := template.New("").ParseFiles("templates/about.html", "templates/base.html")
	err = templ.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}

// function to get the about us page
func GetContactPage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title   string
		Navlink string
	}

	p := Page{
		Title:   "Contact",
		Navlink: "contact",
	}

	templ, err := template.New("").ParseFiles("templates/contact.html", "templates/base.html")
	err = templ.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
