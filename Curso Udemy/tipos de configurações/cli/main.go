// usando flags
package main

import (
	"flag"
	"fmt"
)

func main() {
	// nome da flag | valor default | explicação da flag
	port := flag.String("porta", "8000", "Parametro que se refere a porta para a execução")

	flag.Parse() // é necessário para carregar as flags

	fmt.Println(*port)
}