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
		if r.Method != "POST" { // pode usar também -> http.MethodPost que é uma melhor forma e mais organizada
			w.Header().Set("Allow", "POST") // serve para escrever no Header que o único método permitido é o POST
			// Rejeitar a request
			w.WriteHeader(405) // código para método não permitido
			fmt.Fprint(w, "Método usado é diferente de POST, portanto impossivel continuar!")
			return
		}
		w.Write([]byte("Isso está rodando no /notes/"))
	})

	mux.HandleFunc("/notes/view", func(w http.ResponseWriter, r *http.Request) {
		// segunda forma de negar as requisições
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost) // serve para escrever no Header que o único método permitido é o POST
			http.Error(w, "Método não permitido!", 405) // Erorr escreve a mensagem de erro junto do status code
			// ou
			// http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed) // é a mesma coisa de 405
		}
		w.Write([]byte("Você está vendo alguma nota!"))
	})

	mux.HandleFunc("/notes/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // dizendo no cabeçalho que seŕa uma requisição do tipo json
		fmt.Fprintf(w, `{"id":"teste em json"}`) // eviando o texto em json
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(200)
		w.Write([]byte("Rotas: /notes /notes/view  /notes/create"))
	})
	
	http.ListenAndServe(PORTA, mux)
}