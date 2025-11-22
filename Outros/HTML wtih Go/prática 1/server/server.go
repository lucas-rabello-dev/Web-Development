package server

import (
	"net/http"
)

func ServerStaticFiles(mux *http.ServeMux) *http.ServeMux{
	staticHandler := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	return mux // aparentemente ele retorna o endereço
}

func ServerImageFiles(mux *http.ServeMux) *http.ServeMux{
	imagesHanler := http.FileServer(http.Dir("./images"))
	mux.Handle("/images/", http.StripPrefix("/images/", imagesHanler))
	return mux
}