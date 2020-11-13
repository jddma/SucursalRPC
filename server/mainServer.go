package server

import (
	"fmt"
	"log"

	"../serverUtilities/auth"
)

/**
*	la esta función permite iniciar la ejecución del servidor
*/

func userOptions() {

	for {
		fmt.Print("Presiona enter para registrar un nuevo ATM")
		fmt.Scanln()
		auth.RegisterATM()
	}

}

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
	server.Handle("POST" ,"/controllers/getbalance", server.AddMiddleware(GetBalance, CheckAuth("/"), Log()))
	server.Handle("DELETE" ,"/controllers/deleteaccount", server.AddMiddleware(DeleteAccount, CheckAuth("/"), Log()))
	server.Handle("PUT" ,"/controllers/modifyaccount", server.AddMiddleware(ModifyAccount, CheckAuth("/"), Log()))
	server.Handle("PUT" ,"/controllers/addmoney", server.AddMiddleware(AddMoney, Log()))
	server.Handle("PUT" ,"/controllers/withdrawals", server.AddMiddleware(Withdrawals, CheckAuth("/"), Log()))
	server.Handle("POST" ,"/controllers/addaccount", server.AddMiddleware(AddAccount, CheckAuth("/"), Log()))
	server.Handle("POST" ,"/controllers/login", server.AddMiddleware(Login, Log()))

	//Vistas
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Log()))
	server.Handle("GET", "/panel", server.AddMiddleware(HandlePanel, CheckAuth("/"), Log()))

	go userOptions()

	log.Fatal(server.Listen())
	
}