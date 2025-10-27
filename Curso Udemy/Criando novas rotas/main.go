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
	mux.HandleFunc("/notes/",notes)

	mux.HandleFunc("/notes/view", view)

	mux.HandleFunc("/notes/create", create)

	mux.HandleFunc("/", default_)
	
	http.ListenAndServe(PORTA, mux)
}

// obtendo informações da URL
func view(w http.ResponseWriter, r *http.Request) {
		// segunda forma de negar as requisições
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet) // serve para escrever no Header que o único método permitido é o POST
		http.Error(w, "Método não permitido!", 405) // Error escreve a mensagem de erro junto do status code
		// ou
		// http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed) // é a mesma coisa de 405
		return
	}
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery) // na URL voce pode digitar ?alguma_coisa (precisa do ?) e ele vai printar o alguma_coisa
	w.Write([]byte("Isso é um teste"))
}

func notes(w http.ResponseWriter, r *http.Request) {
		// identificar o metodo da requisição
	if r.Method != "POST" { // pode usar também -> http.MethodPost que é uma melhor forma e mais organizada
		w.Header().Set("Allow", "POST") // serve para escrever no Header que o único método permitido é o POST
		// Rejeitar a request
		w.WriteHeader(405) // código para método não permitido
		fmt.Fprint(w, "Método usado é diferente de POST, portanto impossivel continuar!")
		return
	}
	w.Write([]byte("Isso está rodando no /notes/"))
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // dizendo no cabeçalho que seŕa uma requisição do tipo json
	fmt.Fprintf(w, `{"id":"teste em json"}`) // eviando o texto em json
}


func default_(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Rotas: /notes /notes/view  /notes/create"))
} 