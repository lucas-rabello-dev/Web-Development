package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pratica/server"
)

const PORTA string = ":8000"

var homeHTML string = "./templates/home/home.html"

var buyHTML string = "./templates/home/buy/buy.html"

func main() {

	mux := http.NewServeMux()

	mux = server.ServerStaticFiles(mux)
	
	mux = server.ServerImageFiles(mux)

	mux.HandleFunc("/home/",Home)

	mux.HandleFunc("/home/buy/", Buy)

	fmt.Printf("Running at: http://localhost%s \n", PORTA)

	http.ListenAndServe(PORTA, mux)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método não permitido!",http.StatusMethodNotAllowed)
		log.Fatal("método da requisição não permitido, método usado:", r.Method)
	}

	t, err := template.ParseFiles(homeHTML)
	if err != nil {
		http.Error(w, "Error parse html", http.StatusInternalServerError)
		log.Fatal(err) // for debug in terminal and return 
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "home", nil)

}

func Buy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		log.Fatal("método da requisição não permitido, método usado:", r.Method)
	}

	t, err := template.ParseFiles(buyHTML)
	if err != nil {
		http.Error(w, "Error parse files", http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "buy", nil)
}
