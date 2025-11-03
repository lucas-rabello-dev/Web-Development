package handlers

import (
	"html/template"
	"net/http"
)

const crudPath string = "./templates/crud-html/crud.html"

func CrudHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método HTTP inválido!", http.StatusMethodNotAllowed)
		w.Header().Set("Allow", http.MethodGet)
		return
	}

	t, err := template.ParseFiles(crudPath)
	if err != nil {
		http.Error(w, "Erro ao Carregar .html corretamente!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	t.ExecuteTemplate(w, "crud-html", nil)
}