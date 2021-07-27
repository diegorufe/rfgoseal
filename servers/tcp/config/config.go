package config

import (
	"log"
	"rfgocore/logger"
	"rfgonet/sockets/beans/tcp"
	domainservice "rfgoseal/domain/service/tcp"
)

// Config método para configurar el servidor tcp
//
// @parameter tcpServer es el servidor tcp a configurar
func Config(tcpServer *tcp.TCPServer) {

	// Configuramos el log
	logger.StartLog("rfgoseal.log", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

	// Fijamos la función encargada de procesar los datos recividos
	tcpServer.FunctionProcessDataReceived = domainservice.ProcessTCPClientData
}
