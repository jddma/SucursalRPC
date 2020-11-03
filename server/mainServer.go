package server

import (
	"log"
)

/**
*	la esta función permite iniciar la ejecución del servidor
*/
func RunServer(hostRPCServer string, token string)  {

	startRPCConnection(hostRPCServer, token)

	config := newServerConfig()

	//a la función NewServer() se le debe enviar el puerto a usar
	server := NewServer(config)

	/**
	*	definir en este punto los Handlers correspondientes a las rutas
	*	teniendo en cuante que cada handler solo admite un tipo de solicitud
	*/
		server.Handle("POST" ,"/", server.AddMiddleware(GetBalance, Log()))

	log.Fatal(server.Listen())
	
}