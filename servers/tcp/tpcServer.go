package tcp

import (
	"rfgonet/sockets/beans/tcp"
	"rfgoseal/servers/tcp/config"
)

// StarTCPServer método para lanzar el servidor tcp
func StarTCPServer() {

	var tcpServer *tcp.TCPServer = tcp.NewTCPServer()

	// Cargamos la configuración
	config.Config(tcpServer)

	// lanzamos el servidor
	tcpServer.Listen()
}
