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

// function to get the contact us page
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

// function to get the experience Page
func GetExperiencePage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title   string
		Navlink string
	}

	p := Page{
		Title:   "Experience",
		Navlink: "experience",
	}

	templ, err := template.New("").ParseFiles("templates/experience.html", "templates/base.html")
	err = templ.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}

// function to get the blog Page
func GetBlogPage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title   string
		Navlink string
	}

	p := Page{
		Title:   "Blog",
		Navlink: "blog",
	}

	templ, err := template.New("").ParseFiles("templates/blog.html", "templates/base.html")
	err = templ.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}

// get 404 page
func Custom404Handler(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title   string
		Navlink string
	}

	p := Page{
		Title:   "Error",
		Navlink: "error",
	}

	templ, err := template.New("").ParseFiles("templates/errors/404.html", "templates/base.html")
	err = templ.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
