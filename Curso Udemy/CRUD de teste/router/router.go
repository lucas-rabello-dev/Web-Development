package router

import (
	"net/http"
)

// func para criar as rotas

func Router(mux *http.ServeMux, pattern string, fun func(w http.ResponseWriter, r *http.Request)) *http.ServeMux{

	mux.HandleFunc(pattern, fun)
	// arquivos estáticos (CSS, JS, imagens)

	return mux
}