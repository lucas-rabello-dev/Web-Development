package main

import (
	"fmt"
	"net/http"
)

const PORTA3 string = ":8080" 

type Handler_meu struct {}

func (Handler_meu)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200) // código 200 para informar que tudo deu certo
	w.Write([]byte("Default -- Rota coringa por que terminou com a /"))
}

func handlerEstatico(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Isso é de uma rota estática") // é basicamente um writer universal
}



func main() {
	fmt.Println("Servidor rodando na porta:", PORTA3)
	h := Handler_meu{}

	// tanto Handle quanto HandleFunc tem o DefaultServerMux
	// http.Handle("/", h)

	// rota coringa / ou /hello/
	// rota estática /hello -> aponta diretamente

	// com o mux não precisa de ordem nas rotas
	mux := http.NewServeMux()
	mux.Handle("/", h) // -> rota coringa (default)
	mux.HandleFunc("/hello", handlerEstatico)
	mux.HandleFunc("/hello/world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rodando em /hello/world"))
	})

	// http.ListenAndServe(PORTA3, nil) // -> nil se refere ao server mux

	http.ListenAndServe(PORTA3, mux) // -> você pode passar o mux pois ele tambem tem um método ServeHTTP
}	