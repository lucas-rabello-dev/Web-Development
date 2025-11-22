package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const porta string = ":8000"

const fileHome string = "./interface/home/home.html"

const filePage1 string = "./interface/page1/page1.html"

const filePage2 string = "./interface/page1/page2/page2.html"

func main() {
	mux := http.NewServeMux()	

	// em Dir() podem ser dois caminhos: o absoluto (onde está o arquivo dentro do projeto) e o relativo (normalmente quando tem varios outros styles.css)
	// servindo a pasta do static para que os elementos possam ser acessados
	staticHandler := http.FileServer(http.Dir("./interface/static/"))
	// o primeiro /static/ é o nome que serve o dict que escrevi aqui em cima
	// tipo: a gente nomeou o "./interface/static" de /static/ pique isso
	// o segundo remove esse /static/ para que:
	//URL que chegou: /static/css/css-home/style.css
	// StripPrefix tira o /static/
	// nos usamos o /static/ para acessar o "./interface/static" e ai tiramos por que já estamos servindo ele
	// URL que o FileServer realmente recebe: /css/home/style.css
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))


	mux.HandleFunc("/home/", home)
	mux.HandleFunc("/home/page1", page1)
	mux.HandleFunc("/home/page1/page2", page2)


	fmt.Printf("Rodando em: http://localhost%s/ \n", porta)
	http.ListenAndServe(porta, mux)
}


func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home/"  {
		http.NotFound(w, r) // responde a página como 404
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		fmt.Println("Deu pau (o de cima)")
		return
	}

	t, err := template.ParseFiles(fileHome) // caminho do arquivo html
	if err != nil {
		fmt.Println("Deu pau (o de baixo)", err)
		http.Error(w, "Aconteceu algo de errado!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// o segundo argumento é o nome do template
	t.ExecuteTemplate(w, "home", nil) // o nil é os dados que iriamos alterar comforme o programa Ex: {{ template "blabla" . }}	
}


func page1(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles(filePage1)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na operação parse: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "page1", nil)
}

func page2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles(filePage2)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na operação parse: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "page2", nil)
}