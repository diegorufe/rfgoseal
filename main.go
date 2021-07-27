package main

import (
	"os"
	"rfgoseal/servers/tcp"
)

func main() {
	// Obtenemos los argumentos del ejecutable
	argsWithoutProg := os.Args[1:]

	var action string = "tcp"

	// Obtenemos la acción a ejecutar desde la línea de comandos
	if len(argsWithoutProg) > 0 {
		action = argsWithoutProg[0]
	}

	switch action {

	// Servidor http para manejo de la web
	case "http":
		break

		// Servidor tcp para escuchar mensajes
	case "tcp":
		tcp.StarTCPServer()

	default:
		tcp.StarTCPServer()

	}

}
