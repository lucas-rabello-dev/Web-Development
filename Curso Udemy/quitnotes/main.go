package main

import "net/http"

const PORTA string = ":8080"

// O segundo argumento da função ListenAndServe é uma interface
// que tenha o método ServerHTTP
type Handler struct {}
func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	// |-> só pode ser usado antes do w.Write()

	// Retorna uma resposta
	w.Write([]byte("Deu certo mano!"))
}

func main() {
	handler := Handler{}

	// Primeira forma
	http.ListenAndServe(PORTA, handler) // Ela sobe o server htpp conforme as configurações

	// Segunda forma de criar o servidor

	// pattern -> a fomra de como a URL é acessada
	// http.Handle("/", handler)
	// http.ListenAndServe(PORTA, nil)
}