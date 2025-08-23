package main

import (
	"net/http"
)

const PORTA2 = ":8080" 

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("teste do handler"))
}

func main() {
	http.ListenAndServe(PORTA2, nil)
	http.HandleFunc("/", helloHandler)

}