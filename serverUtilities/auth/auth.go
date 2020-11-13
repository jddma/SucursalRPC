package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

/**
*	la función SetUserSession instanciara una sesión de usuario 
*	válida para su respectiva autenticación, retornara true
*	si el proceso fue correctamente ejecutado false en caso
*	contrario
*/

func OpenConnection() *mongo.Client {

	var result *mongo.Client

	mgdUser := os.Getenv("MGD_USER")
	mgdPassword := os.Getenv("MGD_PASSWORD")
	mgdHost := os.Getenv("MGD_HOST")

	uri := "mongodb://" + mgdUser + ":" + mgdPassword + "@" + mgdHost + ":27017"

	var err error
	result, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = result .Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return result

}

func RegisterATM()  {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type": "atm",
	})
	tokenString, _ := token.SignedString([]byte("KEY"))

	newATM := map[string]string{
		"type": "atm",
		"token": tokenString,
	}

	client := OpenConnection()
	//Persistir los datos en la base de datos
	collection := client.Database("branch").Collection("clients")
	_, err := collection.InsertOne(context.TODO(), newATM)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TOKEN: ", tokenString)

}

func SetUserSession(user UserSession, cookieHandler *securecookie.SecureCookie, w http.ResponseWriter) bool {

	//obtiente los datos necesarios del parametro user
	indentifier := user.GetIdentifier()
	role := user.GetRole()
	
	//establece una estructura para almecenar los datos en la sesión
	value := map[string]string{
		"Name": indentifier,
		"Role": role,
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