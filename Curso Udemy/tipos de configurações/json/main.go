package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// tipo para reŕesentar o conteudo do arquivo json

type Configs struct {
	// struct anonima referente a estrutura o aquivo Json
	Server struct {
		Porta int
		Outro string
	}
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	var configs Configs

	// Encoder -> struct para json
	// Decoder -> json para struc

	err = json.NewDecoder(file).Decode(&configs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Porta: %d | Outro: %s\n", configs.Server.Porta, configs.Server.Outro)


}