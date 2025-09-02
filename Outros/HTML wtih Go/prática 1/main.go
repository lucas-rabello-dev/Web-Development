package main

import (
	"text/template"
	"log"
	"net/http"
)

const PORTA string = ":8080"

var temp_login = template.Must(template.ParseGlob("templates/teste_login/index.html"))
var temp_for = template.Must(template.ParseGlob("templates/teste_for/index.html"))

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/login",LoginRoot)

	mux.HandleFunc("/form", ForRoot)

	log.Println("Running at: http://localhost:8080")

	http.ListenAndServe(PORTA, mux)
}

func LoginRoot(w http.ResponseWriter, r *http.Request) {
	// com o I maiúsculo de propósito 
	temp_login.ExecuteTemplate(w, "Login", nil)
}

func ForRoot(w http.ResponseWriter, r *http.Request) {
	temp_for.ExecuteTemplate(w, "For", nil)
}
