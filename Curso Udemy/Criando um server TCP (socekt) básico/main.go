// No windows você pode fazer a conexão pelo powershell usando a ferramenta telnet
// telnet localhost 5000

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// O primeiro parâmetro é o protocolo e o segundo é a porta que o servidor vai escutar
	listen, err := net.Listen("tcp", ":5000") // Criando o Listen
	if err != nil {
		panic(err)
	}
	// Uma boa prática é sempre fechar os Listen de servidores
	defer listen.Close()

	for {
		// Aceitando a conexão
		con, err := listen.Accept() // Aqui você esta estabelecendo a conexão com o cliente
			if err != nil {
			panic(err)
		}
		fmt.Println("Conexão estabelecida!")
		// O uso dessa gogroutine serve para que quando eu estiver recebendo uma
		// outra requisição o servidor pode escuta-lá sem precisar esperar a requisição
		// atual terminar
		// portando essa função roda ao "mesmo tempo"
		go func(con net.Conn) {
			for {
				// Aqui cria um buffer de leitura lendo a conexão até uma quebra de linha
				data, err := bufio.NewReader(con).ReadString('\n')
				if err != nil {
					panic(err)
				}
				if strings.Contains(data, "exit") {
					break
				}
				fmt.Println("Dado recebido:", data)
							// Enviar uma resposta para o cliente
				con.Write([]byte("Sua mensagem foi recebida com sucesso! \n"))
			}


			con.Close();fmt.Println("A conexão foi encerrada!")
		}(con)


	}
}