package server

import (
	"../serverUtilities/auth"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

/**
*	Para definir un nuevo Middleware puede usar la siguiente plantilla:

	func MiddlewareName() Middleware {
		return func(f http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				f(w, r)
			}
		}
	}	
*
*	tenga en cuenta que la línea f(w, r) es la que permite
*	terminar con la ejecución del middleware actual y continuar con la 	*	*	ejecución del siguiente
*	(en caso de que exista otro)	
*/


/**
*	el middleware CheckAuth permite verificar si un usuario esta autenticado
*	en caso de que no lo este sera redireccinado a la ruta que sea enviada
*	como parametro
*/

type clients struct {
	Token string
	Type string
}
func CheckAuth(URLToRedirect string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			
			redirect := false

			sessionData := auth.GetSessionData(r, cookieHandler)
			_, err1 := sessionData["Name"]
			_, err2 := sessionData["Role"]

			if (! err1) || (! err2){
				redirect = true
			}

      		if redirect{

				connection := auth.OpenConnection()
				jwt := r.Header["Authorization"]
				if jwt == nil {
					http.Redirect(w, r, URLToRedirect, 401)
					return
				}
				filter := bson.D{{
					"token", jwt[0],
				}}
				collection := connection.Database("branch").Collection("clients")
				var queryResult clients
				err := collection.FindOne(context.TODO(), filter).Decode(&queryResult)
				if err != nil {
					http.Redirect(w, r, URLToRedirect, 401)
					return
				}
			}
			f(w, r)
		}

	}

}

//esta middleware registra los logs a los handlers en los que se implementa
func Log() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			defer func() { 

				fmt.Println("-------------------")
				log.Println("\nPATH=", r.URL.Path)
				fmt.Println("HOST=", r.RemoteAddr)
				fmt.Println("Method=", r.Method)
				fmt.Println(time.Since(start))
				fmt.Println("-------------------")

			}()
			f(w, r)

		}

	}

}