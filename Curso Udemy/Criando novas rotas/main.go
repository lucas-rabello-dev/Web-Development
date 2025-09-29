package main

import (
	"fmt"
	"net/http"
)

const PORTA = ":8080"

func main() {
	fmt.Println("Rodando na porta:", PORTA)
	
	mux := http.NewServeMux()

	// HandleFunc pois é mais simples nesse caso
	mux.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request){
		// identificar o metodo da requisição
		if r.Method != "POST" {
			// Rejeitar a request
			w.WriteHeader(405) // código para método não permitido
			fmt.Fprint(w, "Método usado é diferente de POST, portanto impossivel continuar!")
			return
		}
		w.Write([]byte("Isso está rodando no /notes/"))
	})

	mux.HandleFunc("/notes/view", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Você está vendo alguma nota!"))
	})

	mux.HandleFunc("/notes/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Você está criando uma nota!"))
	})
	
	http.ListenAndServe(PORTA, mux)
}