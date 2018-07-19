package main

import (
	"io"
	"net/http"
)

func none(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "You're at the home page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I'm a fricken dog, dude")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Ivan")
}

func main() {
    http.Handle("/me/", http.HandlerFunc(me))
    http.Handle("/dog/", http.HandlerFunc(dog))
    http.Handle("/", http.HandlerFunc(none))

    http.ListenAndServe(":8080", nil)
}