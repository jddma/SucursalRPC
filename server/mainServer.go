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

	//Controladores para los formularios
	server.Handle("POST" ,"/controllers/getbalance", server.AddMiddleware(GetBalance, Log()))
	server.Handle("DELETE" ,"/controllers/deleteaccount", server.AddMiddleware(DeleteAccount, Log()))
	server.Handle("PUT" ,"/controllers/modifyaccount", server.AddMiddleware(ModifyAccount, Log()))
	server.Handle("PUT" ,"/controllers/addmoney", server.AddMiddleware(AddMoney, Log()))
	server.Handle("PUT" ,"/controllers/withdrawals", server.AddMiddleware(Withdrawals, Log()))
	server.Handle("POST" ,"/controllers/addaccount", server.AddMiddleware(AddAccount, Log()))
	server.Handle("POST" ,"/controllers/login", server.AddMiddleware(Login, Log()))

	//Vistas
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Log()))
	server.Handle("GET", "/panel", server.AddMiddleware(HandlePanel, CheckAuth("/"), Log()))

	log.Fatal(server.Listen())
	
}