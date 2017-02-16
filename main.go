package main

import (
	"github.com/alecthomas/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var tpl *template.Template

type Person struct {
	Name string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/about", About)
	router.GET("/user/:name", User)

	http.ListenAndServe(":8080", router)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func User(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Agent := Person{
		p.ByName("name"),
	}

	tpl.ExecuteTemplate(w, "user.gohtml", Agent)
}
