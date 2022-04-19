package main

import (
	"keyvaluepairexercise/httpServer"
	"keyvaluepairexercise/keyValueService"
	"keyvaluepairexercise/tcpServer"
)

func main() {
	keyValueService.Initialise()

	go tcpServer.TcpListen()
	httpServer.HttpListen()
}
