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
	server.Handle("POST" ,"/controllers/getbalance", server.AddMiddleware(GetBalance, Log()))
	server.Handle("DELETE" ,"/controllers/deleteaccount", server.AddMiddleware(DeleteAccount, Log()))
	server.Handle("PUT" ,"/controllers/modifyaccount", server.AddMiddleware(ModifyAccount, Log()))


	log.Fatal(server.Listen())
	
}