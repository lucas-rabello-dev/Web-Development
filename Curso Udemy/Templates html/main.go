package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const porta string = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", Home)
	mux.HandleFunc("/home/page2", Page2)

	fmt.Println("Rodando na porta: http://localhost" + porta)
	http.ListenAndServe(porta, mux)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Alllow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed) // 405
		return 
	}
	w.WriteHeader(http.StatusOK) // 200

	t, err := template.ParseFiles("./interface/home/home.html")
	if err != nil {
		http.Error(w, "Aconteceu algum erro!", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

func Page2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed) // 405
		return 
	}
	w.WriteHeader(http.StatusOK) // 200

	t, err := template.ParseFiles("./interface/home/page2/page2.html")
	if err != nil {
		http.Error(w, "Aconteceu algum erro!", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, t.Name(), nil)

}
