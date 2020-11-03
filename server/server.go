package server

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

//struct Server ...
type Server struct{

	config *serverConfig
	router *Router

}

/**
*	variable para definir las cookies de sesión
*
*	---NO MODIFICARLA EN NINGUN PUNTO DE EJECUCIÓN---
*/
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

//función encargada de implementar los middlewares deseados a un handler
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		f = m(f)
	}
	return f

}


//función que se ejecutara para iniciar el servidor
func (s *Server) Listen() error{

	http.Handle("/", s.router)

	//estas dos líneas sirven el directorio que contiene los archivos estaticos a usar
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	var err error
	if s.config.Ssl{
		err = http.ListenAndServeTLS(s.config.Port, s.config.SslCertPath, s.config.SslKeyPath, nil)
	}else {
		err = http.ListenAndServe(s.config.Port, nil)
	}

	if err != nil{
		fmt.Println("err")
		return err
	}
	return nil

}

//función usada para definir los Handles del servidor con sus respectias rutas
func (s *Server) Handle(method string, path string, handler http.HandlerFunc){

	_, exist := s.router.rules[path]
	if !exist{
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler

}

//NewServer ...
func NewServer(config *serverConfig) *Server{

	return &Server{
		config: config,
		router: NewRouter(),
	}

}

