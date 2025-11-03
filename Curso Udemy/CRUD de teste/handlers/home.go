package handlers

// onde ficam as funções de resposta para as rotas

import (
	"html/template"
	"net/http"
)

const homePath string = "./templates/home/home.html"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método HTTP inválido!", http.StatusMethodNotAllowed)
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	t, err := template.ParseFiles(homePath)
	if err != nil {
		http.Error(w, "Erro ao Carregar .html corretamente!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	t.ExecuteTemplate(w, "home", nil)
}	