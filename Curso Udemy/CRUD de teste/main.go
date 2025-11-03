package main

import (
	"crud-teste/handlers"
	"crud-teste/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux = router.Router(mux, "/home", handlers.Home)
	mux = router.Router(mux, "/crud", handlers.CrudHandler)

	fmt.Printf("Rodando em: http.localhost:8000/home \n")
	log.Fatal(http.ListenAndServe(":8000", mux))
}