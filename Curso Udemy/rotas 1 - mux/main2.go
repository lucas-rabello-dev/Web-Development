package main

import (
	"fmt"
	"net/http"
)

const PORTA2 = ":8080" 

type HandlerTeste struct {}

func (HandlerTeste)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("teste do handler"))
}

func main() {
	fmt.Println("Rodando na porta:", PORTA2)
	
	// http.HandleFunc("/", helloHandler)
	// ou você pode também fazer:
	h := HandlerTeste{} // Instancia a struct pois ela é uma interface necessária
	http.Handle("/hello", h)
	http.ListenAndServe(PORTA2, nil) // sempre por último
}