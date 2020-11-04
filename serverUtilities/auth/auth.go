package auth

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

/**
*	la función SetUserSession instanciara una sesión de usuario 
*	válida para su respectiva autenticación, retornara true
*	si el proceso fue correctamente ejecutado false en caso
*	contrario
*/
func SetUserSession(user UserSession, cookieHandler *securecookie.SecureCookie, w http.ResponseWriter) bool {

	//obtiente los datos necesarios del parametro user
	indentifier := user.GetIdentifier()
	role := user.GetRole()
	
	//establece una estructura para almecenar los datos en la sesión
	value := map[string]string{
		"name": indentifier,
		"role": role,
	}

	//codificar la cookie
	encoded, err := cookieHandler.Encode("session", value)

	if  err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)

		return true
	}

	return false

}

/**
*	la función CleanUserSession limpiara la cookie session
*	esto sera para terminar la sesión de usuario
*/
func CleanUserSession(r http.ResponseWriter) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(r, cookie)

}

/**
*	la función GetSessionData retornara una estructura que contendra
*	la información guardada en la sesión (en caso de que exista)
*/
func GetSessionData(r *http.Request, cookieHandler *securecookie.SecureCookie) map[string]string {

	var dataSession map[string]string
	cookie, err := r.Cookie("session") 
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			dataSession = cookieValue
		}
	}
	return dataSession

}

/**
*	la intefase Usersession es un clon de la interface contenida en 
*	en el paquete server y cumple la misma función
*/
type UserSession interface{

	GetIdentifier() string
	GetPassword() string
	GetRole() string

}