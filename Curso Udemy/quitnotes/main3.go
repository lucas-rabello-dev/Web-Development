package main

import (
	"fmt"
	"net/http"
)

const PORTA3 string = ":8080" 

type Handler_meu struct {}

func (Handler_meu)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200) // código 200 para informar que tudo deu certo
	w.Write([]byte("Isso esta rodando mano"))
}

func main() {
	fmt.Println("Servidor rodando na porta:", PORTA3)
	h := Handler_meu{}

	// tanto Handle quanto HandleFunc tem o DefaultServerMux
	// http.Handle("/", h)

	mux := http.NewServeMux()
	mux.Handle("/hello", h)

	// http.ListenAndServe(PORTA3, nil) // -> nil se refere ao server mux

	http.ListenAndServe(PORTA3, mux) // -> você pode passar o mux pois ele tambem tem um método ServeHTTP
}	