package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const porta string = ":8000"

const fileHome string = "./interface/home/home.html"

const filePage1 string = "./interface/page1/page1.html"

const filePage2 string = "./interface/page1/page2/page2.html"

func main() {
	mux := http.NewServeMux()	


	mux.HandleFunc("/home", home)
	mux.HandleFunc("/home/page1", page1)
	mux.HandleFunc("/home/page1/page2", page2)


	fmt.Printf("Rodando em: http://localhost%s/ \n", porta)
	http.ListenAndServe(porta, mux)
}


func home(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		fmt.Println("Deu pau (o de cima)")
		return
	}

	t, err := template.ParseFiles(fileHome) // caminho do arquivo html
	if err != nil {
		fmt.Println("Deu pau (o de baixo)", err)
		http.Error(w, "Aconteceu algo de errado!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// o segundo argumento é o nome do template
	t.ExecuteTemplate(w, "home", nil) // o nil é os dados que iriamos alterar comforme o programa Ex: {{ template "blabla" . }}	
}


func page1(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles(filePage1)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na operação parse: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "page1", nil)
}

func page2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles(filePage2)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na operação parse: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "page2", nil)
}