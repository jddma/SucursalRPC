package main

import (
	"./server"
	"fmt"
)

func main (){

	var hostRPCServer string
	fmt.Print("Digite el host del servidor de RPC: ")
	fmt.Scanln(&hostRPCServer)

	var token string
	fmt.Print("Digite el token de autorización del servidor de RPC: ")
	fmt.Scanln(&token)

	server.RunServer(hostRPCServer, token)

}