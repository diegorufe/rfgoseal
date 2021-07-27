package tcp

import (
	"encoding/json"
	"net"
	"rfgocore/logger"
	"rfgocore/utils/utilsstring"
	"rfgonet/sockets/beans/tcp"
	"rfgoseal/domain/core/constants"
)

// ProcessTCPClientData método encargado de procesar los datos recividos en el servidor tcp
//
// @parameter connection es la conexión del cliente
//
// @parameter tcpServer es el servidor tcp
//
// @parameter data son los datos recividos del cliente
func ProcessTCPClientData(connection net.Conn, tcpServer *tcp.TCPServer, data []byte) {

	// if logger.DefaultLogger().IsDebugEnabled() && len(data) > 0 {
	// 	logger.DefaultLogger().Debug(utilsstring.TrimAllSpace(string(data)))
	// }

	jsonData := utilsstring.TrimAllSpace(string(data))
	var mapLogData map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &mapLogData)

	if err != nil {
		if logger.DefaultLogger().IsErrorEnabled() {
			logger.DefaultLogger().Error(err)
		}
	} else {
		// if logger.DefaultLogger().IsDebugEnabled() {
		// 	logger.DefaultLogger().Debug(result)
		// }

		if application, okApp := mapLogData[constants.ParamLogDataApplication]; okApp {
			if logger.DefaultLogger().IsDebugEnabled() {
				logger.DefaultLogger().Debug(application)
			}

			// TODO buscar collection de la aplicación e insetar log de la aplicación
		}

	}

}
