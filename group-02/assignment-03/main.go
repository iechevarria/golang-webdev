package main

import (
	"net/http"
	"html/template"
)

type page struct {
	Name string
	Content string
}

func none(w http.ResponseWriter, r *http.Request) {
	p := page{"Home", "You're on the home page"}
	tpl.ExecuteTemplate(w, "tpl.gohtml", p)
}

func dog(w http.ResponseWriter, r *http.Request) {
	p := page{"Dog", "I'm a fricken dog, dude"}
	tpl.ExecuteTemplate(w, "tpl.gohtml", p)
}

func me(w http.ResponseWriter, r *http.Request) {
	p := page{"Me", "Ivan"}
	tpl.ExecuteTemplate(w, "tpl.gohtml", p)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    http.Handle("/me/", http.HandlerFunc(me))
    http.Handle("/dog/", http.HandlerFunc(dog))
    http.Handle("/", http.HandlerFunc(none))

    http.ListenAndServe(":8080", nil)
}