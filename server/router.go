package server

import (
	"net/http"
)


//sruct Router
type Router struct{

	rules map[string]map[string]http.HandlerFunc

}

//función encargada de recibir las solicitudes hechas al servidor
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	//llama a la función que busca el handler deacuerdo a la URL de la solicitud
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)

	//Verificar si la URL solicitada existe
	if !exist{
		w.WriteHeader(http.StatusNotFound)
		return
	}	

	//verifica que el metodo este permitido para la ruta
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}

//función que se encarga de buscar el handler solicitado
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {

	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

//funcíon para instanciar el router
func NewRouter() *Router{

	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}

}