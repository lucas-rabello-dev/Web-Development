package main

import "net/http"

const PORTA string = ":8080"

// O segundo argumento da função ListenAndServe é uma interface
// que tenha o método ServerHTTP
type Handler_hello struct {}
func (Handler_hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	// |-> só pode ser usado antes do w.Write()

	// Retorna uma resposta
	w.Write([]byte("handler do hello"))
}

type Handler_world struct {}
func (Handler_world) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handler do world"))
}

type Handler struct {}
func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handler normal"))
}

func main() {
	handler_hello := Handler_hello{}
	Handler_world := Handler_world{}
	handler := Handler{}

	// Primeira forma
	// http.ListenAndServe(PORTA, handler) // Ela sobe o server htpp conforme as configurações

	// Segunda forma de criar o servidor

	// pattern -> a forma de como a URL é acessada
	http.Handle("/hello", handler_hello)
	http.Handle("/world", Handler_world)
	http.Handle("/", handler)
	
	http.ListenAndServe(PORTA, nil)
}