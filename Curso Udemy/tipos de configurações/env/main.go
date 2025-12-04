package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load("")
	if err != nil {
		log.Fatal(err)
	}
	// utilizamos o Getenv pois a função .Load() já carregou as váriaveis de ambiente
	// use Getenv para variáveis opicionais
	// e use o LookupEnv para coisas importantes (portas, IDs etc)
	port := os.Getenv("PORTA")
	fmt.Println(port)
}

func opcao1() {
	// método mais simples
	// retorna uma string 
	// result := os.Getenv("PORTA")
	// se a váriavel de ambiente não existir ele retorna uma string vazia e se ela existir mas não tiver nada lá também retorna vazio
	// fmt.Printf("Server rodando na porta: %s \n", result)


	// metodo que eu vou usar

	result, verificar := os.LookupEnv("PORTA")
	if !verificar {
		log.Fatal("Houve um erro em carregar o valor de .env")
		// podemos também definir uma porta default nesse caso
		// EX:
		// result = "8000"
		// ou esse IF
		// if !verificar || result == "" { result = "8000" }
	}

	fmt.Println(result)
}
