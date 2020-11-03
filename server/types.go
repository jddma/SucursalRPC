package server

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

/**
*la interfase UserSession define los metodos get que debe implementar
*una estructura que se quiera guardar en la sesión
*/
type UserSession interface{

	GetId() int
	GetIdentifier() string
	GetPassword() string

}

//el contexto por defecto del index
type IndexContext struct{
	Title string
}